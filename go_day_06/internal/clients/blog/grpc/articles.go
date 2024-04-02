package grpcclient

import (
	"articles/internal/code"
	"articles/pkg/logger/sl"
	blog "articles/protos/gen/go/articles"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
)

type Article struct {
	Id    int64
	Title string
	Text  string
}

type ArticleData struct {
	Articles []Article
	PrevPage int
	NextPage int
	LastPage int
	Page     int
}

func (c *Client) CreateArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "internal/clients/blog/CreateArticle"
		if r.Method == http.MethodGet {
			claims, err := ParseToken(c.Token, c.Log)
			if err != nil {
				http.Redirect(w, r, "/access_denied", http.StatusSeeOther)
				return
			}

			isAdmin, err := c.Auth.IsAdmin(context.Background(), &blog.IsAdminRequest{
				UserId: int64(claims["uid"].(float64)),
			})
			if err != nil {
				return
			}

			if isAdmin.IsAdmin {
				ts, err := template.ParseFiles("./static/post_article.html")
				if err != nil {
					c.Log.Error("cannot parse page", op, sl.Err(err))
					return
				}
				ts.Execute(w, "")
			} else {
				http.Redirect(w, r, "/articles/my/?page=1", http.StatusSeeOther)
				return
			}
		} else if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				c.Log.Error("failed to get parameters", op, sl.Err(err))
				return
			}

			title := r.PostForm.Get("title")
			text := r.PostForm.Get("text")

			resp, err := c.Articles.CreateArticle(context.Background(), &blog.ArticleData{
				Title: title,
				Text:  text,
			})

			if err != nil {
				if errors.Is(err, code.TitleIsRequired) || errors.Is(err, code.TextIsRequired) {
					c.Log.Warn("title or text is required", op, sl.Err(err))
				} else if errors.Is(err, code.ArticleAlreadyExists) {
					c.Log.Warn("article already exists", op, sl.Err(err))
				} else {
					c.Log.Error("failed to create new article", err)
				}
			} else {
				c.Log.Info("new article created", slog.Int64("article_id", resp.Id))
			}
			http.Redirect(w, r, "/articles/my/?page=1", http.StatusSeeOther)
		} else {
			c.Log.Error("method is not allowed", slog.String("method", r.Method))
		}
	}
}

func (c *Client) ShowMyArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "internal/clients/blog/ShowArticle"

		if r.Method == http.MethodGet {
			id, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil {
				c.Log.Error("failed to get article id", op, sl.Err(err))
				http.Redirect(w, r, "/access_denied", http.StatusSeeOther)
			}

			claims, err := ParseToken(c.Token, c.Log)
			if err != nil {
				http.Redirect(w, r, "/access_denied", http.StatusSeeOther)
				return
			}

			resp, err := c.Articles.ShowMyArticle(context.Background(), &blog.ArticleId{
				Id: int64(id),
			})

			if err != nil {
				if errors.Is(err, code.ArticleIDIsRequired) {
					c.Log.Warn("failed to get article id", op, sl.Err(err))
				} else if errors.Is(err, code.ArticleNotFound) {
					c.Log.Warn("article not found", op, sl.Err(err))
				} else {
					c.Log.Error("failed to get article", err)
				}
				http.Redirect(w, r, "/access_denied", http.StatusSeeOther)
				return
			} else if resp.Author != int64(claims["uid"].(float64)) {
				c.Log.Error("access denied", slog.Int64("user_id", resp.Author))
				http.Redirect(w, r, "/access_denied", http.StatusSeeOther)
				return
			} else {
				c.Log.Info("article received", slog.String("title", resp.Title))
			}

			ts, err := template.ParseFiles("./static/article.html")
			if err != nil {
				c.Log.Error("cannot parse page", op, sl.Err(err))
				return
			}

			ts.Execute(w, Article{
				Title: resp.GetTitle(),
				Text:  resp.GetText(),
			})
		} else {
			c.Log.Error("method is not allowed", slog.String("method", r.Method))
		}
	}
}

func (c *Client) ShowAllMyArticles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "internal/clients/blog/ShowAllArticles"

		if r.Method == http.MethodGet {
			claims, err := ParseToken(c.Token, c.Log)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			resp, err := c.Articles.ShowAllMyArticles(context.Background(), &blog.ArticlesRequest{
				Id: int64(claims["uid"].(float64)),
			})

			if err != nil {
				c.Log.Error("failed to get articles", op, sl.Err(err), err.Error())
				http.Redirect(w, r, "/articles/my/?page=1", http.StatusSeeOther)
				return
			}

			page, err := strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				c.Log.Error("failed to get page number", op, sl.Err(err))
				http.Redirect(w, r, "/articles/my/?page=1", http.StatusSeeOther)
				return
			}

			articles := CreateArticlesList(resp.GetId(), resp.GetTitles(), resp.GetTexts(), page)

			ts, err := template.ParseGlob("./static/*.html")
			if err != nil {
				c.Log.Error("cannot parse page", op, sl.Err(err))
				return
			}

			c.Log.Info("articles received", slog.String("articles", resp.String()))
			ts.ExecuteTemplate(w, "my_articles.html", articles)
		} else {
			c.Log.Error("method is not allowed", slog.String("method", r.Method))
		}
	}
}

func (c *Client) ShowAllArticles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "internal/clients/blog/ShowAllArticles"

		if r.Method == http.MethodGet {
			ts, err := template.ParseFiles("./static/all_articles.html")
			if err != nil {
				c.Log.Error("cannot parse page", op, sl.Err(err))
				return
			}

			resp, err := c.Articles.ShowAllArticles(context.Background(), &emptypb.Empty{})

			if err != nil {
				c.Log.Error("failed to get articles", op, sl.Err(err), err.Error())
				http.Redirect(w, r, "/access_denied", http.StatusSeeOther)
				return
			}

			type Data struct {
				Articles []Article
			}

			page, err := strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				c.Log.Error("failed to get page number", op, sl.Err(err))
				http.Redirect(w, r, "/articles/my/?page=1", http.StatusSeeOther)
				return
			}

			articles := CreateArticlesList(resp.GetId(), resp.GetTitles(), resp.GetTexts(), page)

			c.Log.Info("articles received", slog.String("articles", resp.String()))
			ts.Execute(w, articles)
		} else {
			c.Log.Error("method is not allowed", slog.String("method", r.Method))
		}
	}
}

func CreateArticlesList(id []int64, titles, texts []string, page int) ArticleData {
	var articles []Article
	for i := 0; i < len(id); i++ {
		if len(texts[i]) > 300 {
			texts[i] = texts[i][:300]
		}
		articles = append(articles, Article{
			Id:    id[i],
			Title: titles[i],
			Text:  texts[i] + "...",
		})
	}

	var last int
	if len(articles)%3 == 0 {
		last = len(articles) / 3
	} else {
		last = len(articles)/3 + 1
	}

	begin := page*3 - 3
	end := page * 3
	if end > len(articles)-1 {
		end = len(articles)
	}
	data := ArticleData{
		Articles: articles[begin:end],
		PrevPage: page - 1,
		NextPage: page + 1,
		Page:     page,
		LastPage: last,
	}
	return data
}

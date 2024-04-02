package articles

import (
	"articles/internal/code"
	"articles/internal/domain/models"
	"articles/internal/grpc/auth"
	"articles/internal/services/articles"
	"articles/protos/gen/go/articles"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerArticle struct {
	blog.UnimplementedArticlesServer
	articles Articles
}

type Articles interface {
	CreateArticle(ctx context.Context, title, text string, author int64) (int64, error)
	GetArticle(ctx context.Context, id int64) (string, string, int64, error)
	ShowAllMyArticles(ctx context.Context, id int64) ([]models.Article, error)
	ShowAllArticles(ctx context.Context) ([]models.Article, error)
}

func Register(s *grpc.Server, articles Articles) {
	blog.RegisterArticlesServer(s, &ServerArticle{articles: articles})
}

func (s *ServerArticle) CreateArticle(ctx context.Context, req *blog.ArticleData) (*blog.ArticleId, error) {
	if err := ValidateArticle(req); err != nil {
		return nil, err
	}

	id, err := s.articles.CreateArticle(ctx, req.Title, req.Text, req.Author)
	if err != nil {
		if errors.Is(err, articles.ErrArticleExists) {
			return nil, code.ArticleAlreadyExists
		}
		return nil, code.InternalError
	}

	return &blog.ArticleId{
		Id: id,
	}, nil
}

func ValidateArticle(req *blog.ArticleData) error {
	if req.GetTitle() == "" {
		return code.TitleIsRequired
	}

	if req.GetText() == "" {
		return code.TextIsRequired
	}

	return nil
}

func (s *ServerArticle) ShowMyArticle(ctx context.Context, req *blog.ArticleId) (*blog.ArticleData, error) {
	if err := ValidateId(req); err != nil {
		return nil, err
	}

	title, text, id, err := s.articles.GetArticle(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, articles.ErrArticleNotFound) {
			return nil, code.ArticleNotFound
		}
		return nil, code.InternalError
	}

	return &blog.ArticleData{
		Title:  title,
		Text:   text,
		Author: id,
	}, nil
}

func ValidateId(req *blog.ArticleId) error {
	if req.GetId() == auth.EmptyValue {
		return code.ArticleIDIsRequired
	}

	return nil
}

func (s *ServerArticle) ShowAllMyArticles(ctx context.Context, req *blog.ArticlesRequest) (*blog.ArticlesResponse, error) {
	resp, err := s.articles.ShowAllMyArticles(ctx, req.GetId())
	if err != nil {
		return nil, code.InternalError
	}

	return s.ReturnResponse(resp), nil
}

func (s *ServerArticle) ShowAllArticles(ctx context.Context, _ *emptypb.Empty) (*blog.ArticlesResponse, error) {
	resp, err := s.articles.ShowAllArticles(ctx)
	if err != nil {
		return nil, code.InternalError
	}

	return s.ReturnResponse(resp), nil
}

func (s *ServerArticle) ReturnResponse(resp []models.Article) *blog.ArticlesResponse {
	var ids []int64
	var titles []string
	var texts []string
	for _, v := range resp {
		ids = append(ids, v.ID)
		titles = append(titles, v.Title)
		texts = append(texts, v.Text)
	}
	return &blog.ArticlesResponse{
		Id:     ids,
		Titles: titles,
		Texts:  texts,
	}
}

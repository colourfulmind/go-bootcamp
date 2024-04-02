package articles

import (
	"articles/internal/domain/models"
	"articles/internal/storage"
	"articles/pkg/logger/sl"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

type Article struct {
	Log             *slog.Logger
	ArticleSaver    ArticleSaver
	ArticleProvider ArticleProvider
	TokenTTL        time.Duration
}

type ArticleSaver interface {
	SaveArticle(ctx context.Context, title, text string, author int64) (int64, error)
}

type ArticleProvider interface {
	GetArticle(ctx context.Context, id int64) (models.Article, error)
	GetAllMyArticles(ctx context.Context, id int64) ([]models.Article, error)
	GetAllArticles(ctx context.Context) ([]models.Article, error)
}

var (
	ErrArticleExists   = errors.New("article already exists")
	ErrArticleNotFound = errors.New("article does not found")
)

func New(log *slog.Logger, articleSaver ArticleSaver, articleProvider ArticleProvider, tokenTTL time.Duration) *Article {
	return &Article{
		Log:             log,
		ArticleSaver:    articleSaver,
		ArticleProvider: articleProvider,
		TokenTTL:        tokenTTL,
	}
}

func (a *Article) CreateArticle(ctx context.Context, title, text string, author int64) (int64, error) {
	const op = "internal/services/articles/CreateArticle"
	log := a.Log.With(slog.String("op", op), slog.String("title", title))
	log.Info("attempting to save article")

	id, err := a.ArticleSaver.SaveArticle(ctx, title, text, author)
	if err != nil {
		if errors.Is(err, storage.ErrArticleExists) {
			log.Warn("article already exists")
			return 0, fmt.Errorf("%s: %w", op, ErrArticleExists)
		}
		log.Error("failed to save article", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("article is saved")
	return id, nil
}

func (a *Article) GetArticle(ctx context.Context, id int64) (string, string, int64, error) {
	const op = "internal/services/articles/GetArticle"
	log := a.Log.With(slog.String("op", op), slog.Int64("article id", id))
	log.Info("attempting to get article")

	article, err := a.ArticleProvider.GetArticle(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrArticleExists) {
			log.Warn("article does not exist", sl.Err(err))
			return "", "", 0, fmt.Errorf("%s: %w", op, ErrArticleExists)
		}
		log.Error("failed to get article", sl.Err(err))
		return "", "", 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("successfully got article")
	return article.Title, article.Text, article.Author, nil
}

func (a *Article) ShowAllMyArticles(ctx context.Context, id int64) ([]models.Article, error) {
	const op = "internal/services/articles/ShowAllMyArticles"
	log := a.Log.With(slog.String("op", op))
	log.Info("attempting to get all articles")

	articles, err := a.ArticleProvider.GetAllMyArticles(ctx, id)
	if err != nil {
		log.Error("failed to get articles", sl.Err(err))
		return []models.Article{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("successfully got articles")
	return articles, nil
}

func (a *Article) ShowAllArticles(ctx context.Context) ([]models.Article, error) {
	const op = "internal/services/articles/ShowAllArticles"
	log := a.Log.With(slog.String("op", op))
	log.Info("attempting to get all articles")

	articles, err := a.ArticleProvider.GetAllArticles(ctx)
	if err != nil {
		log.Error("failed to get articles", sl.Err(err))
		return []models.Article{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("successfully got articles")
	return articles, nil
}

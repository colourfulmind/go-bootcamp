package gormdb

import (
	"articles/internal/domain/models"
	"articles/internal/storage"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

type Storage struct {
	DB *gorm.DB
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}

func New(p Postgres) (*Storage, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		p.Host, p.Port, p.User, p.DBName, p.Password, p.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Article{})
	if err != nil {
		return nil, err
	}

	return &Storage{
		DB: db,
	}, nil
}

func (db *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	var user models.User
	check := db.DB.WithContext(ctx).Find(&user, "email = ?", email)
	if check.RowsAffected != 0 {
		return 0, storage.ErrUserExists
	}

	var users []models.User
	ids := db.DB.WithContext(ctx).Find(&users)
	id := ids.RowsAffected

	db.DB.Create(&models.User{
		ID:      id + 1,
		Email:   email,
		PassHas: passHash,
	})

	return id + 1, nil
}

func (db *Storage) User(ctx context.Context, email string) (models.User, error) {
	var user models.User
	db.DB.WithContext(ctx).Where(&models.User{Email: email}).First(&user)

	if user.ID == 0 {
		return models.User{}, storage.ErrUserNotFound
	}

	return models.User{
		ID:      user.ID,
		Email:   user.Email,
		PassHas: user.PassHas,
	}, nil
}

func (db *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	var user models.User
	db.DB.WithContext(ctx).Where(&models.User{ID: userID}).First(&user)

	if user.ID == 0 {
		return false, storage.ErrUserNotFound
	}

	if user.IsAdmin {
		return true, nil
	}

	return false, nil
}

func (db *Storage) MakeAdmin(ctx context.Context, userID int64) error {
	var user models.User
	db.DB.WithContext(ctx).Where(&models.User{ID: userID}).First(&user)

	if user.ID == 0 {
		return storage.ErrUserNotFound
	}

	user.IsAdmin = true
	db.DB.WithContext(ctx).Save(&user)

	return nil
}

func (db *Storage) SaveArticle(ctx context.Context, title, text string, userID int64) (int64, error) {
	var article models.Article
	check := db.DB.WithContext(ctx).Find(&article, "title = ?", title)
	if check.RowsAffected != 0 {
		return 0, storage.ErrArticleExists
	}

	var articles []models.Article
	ids := db.DB.WithContext(ctx).Find(&articles)
	id := ids.RowsAffected

	db.DB.Create(&models.Article{
		ID:     id + 1,
		Author: userID,
		Title:  title,
		Text:   text,
	})

	return id + 1, nil
}

func (db *Storage) GetArticle(ctx context.Context, id int64) (models.Article, error) {
	var article models.Article
	db.DB.WithContext(ctx).Where(&models.Article{ID: id}).First(&article)

	if article.ID == 0 {
		return models.Article{}, storage.ErrArticleNotFound
	}

	return article, nil
}

func (db *Storage) GetAllMyArticles(ctx context.Context, id int64) ([]models.Article, error) {
	var articles []models.Article
	db.DB.WithContext(ctx).Table("articles").Select("id", "title", "text").Where("author = ?",
		strconv.Itoa(int(id))).Scan(&articles)

	return articles, nil
}

func (db *Storage) GetAllArticles(ctx context.Context) ([]models.Article, error) {
	var articles []models.Article
	db.DB.WithContext(ctx).Table("articles").Select("id", "title", "text").Scan(&articles)

	return articles, nil
}

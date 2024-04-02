package storage

import "errors"

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")

	ErrArticleExists   = errors.New("article already exists")
	ErrArticleNotFound = errors.New("article does not exist")
)

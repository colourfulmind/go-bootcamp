package ewrap

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	UserAlreadyExists      = status.Error(codes.AlreadyExists, "user already exists")
	UserNotFound           = status.Error(codes.NotFound, "user not found")
	UserIDIsRequired       = status.Error(codes.InvalidArgument, "user_id is requires")
	EmailIsRequired        = status.Error(codes.InvalidArgument, "email is required")
	PasswordIsRequired     = status.Error(codes.InvalidArgument, "password is required")
	InvalidEmailOrPassword = status.Error(codes.InvalidArgument, "invalid email or password")
	InternalError          = status.Error(codes.Internal, "internal error")

	ArticleAlreadyExists = status.Error(codes.AlreadyExists, "article with the same title already exists")
	ArticleNotFound      = status.Error(codes.NotFound, "article not found")
	ArticleIDIsRequired  = status.Error(codes.InvalidArgument, "id is required")
	TitleIsRequired      = status.Error(codes.InvalidArgument, "title is required")
	TextIsRequired       = status.Error(codes.InvalidArgument, "text is required")
)

package auth

import (
	"articles/internal/ewrap"
	"articles/internal/services/auth"
	"articles/internal/storage"
	"articles/protos/gen/go/articles"
	"context"
	"errors"
	"google.golang.org/grpc"
)

type ServerAuth struct {
	blog.UnimplementedAuthServer
	auth Auth
}

type Auth interface {
	Login(ctx context.Context, email string, password string) (string, error)
	RegisterNewUser(ctx context.Context, email string, password string) (int64, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

func Register(s *grpc.Server, auth Auth) {
	blog.RegisterAuthServer(s, &ServerAuth{auth: auth})
}

const EmptyValue = 0

func (s *ServerAuth) RegisterNewUser(ctx context.Context, req *blog.RegisterRequest) (*blog.RegisterResponse, error) {
	if err := ValidateRegister(req); err != nil {
		return nil, err
	}

	userID, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, auth.ErrUserExists) {
			return nil, ewrap.UserAlreadyExists
		}
		return nil, ewrap.InternalError
	}

	return &blog.RegisterResponse{
		UserId: userID,
	}, nil
}

func ValidateRegister(req *blog.RegisterRequest) error {
	if req.GetEmail() == "" {
		return ewrap.EmailIsRequired
	}
	if req.GetPassword() == "" {
		return ewrap.PasswordIsRequired
	}

	return nil
}

func (s *ServerAuth) Login(ctx context.Context, req *blog.LoginRequest) (*blog.LoginResponse, error) {
	if err := ValidateLogin(req); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, ewrap.InvalidEmailOrPassword
		}
		return nil, ewrap.InternalError
	}

	return &blog.LoginResponse{
		Token: token,
	}, nil
}

func ValidateLogin(req *blog.LoginRequest) error {
	if req.GetEmail() == "" {
		return ewrap.EmailIsRequired
	}

	if req.GetPassword() == "" {
		return ewrap.PasswordIsRequired
	}

	return nil
}

func (s *ServerAuth) IsAdmin(ctx context.Context, req *blog.IsAdminRequest) (*blog.IsAdminResponse, error) {
	if err := ValidateAdmin(req); err != nil {
		return nil, err
	}

	isAdmin, err := s.auth.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return nil, ewrap.UserNotFound
		}
		return nil, ewrap.InternalError
	}

	return &blog.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}

func ValidateAdmin(req *blog.IsAdminRequest) error {
	if req.GetUserId() == EmptyValue {
		return ewrap.UserIDIsRequired
	}

	return nil
}

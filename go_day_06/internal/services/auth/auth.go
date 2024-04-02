package auth

import (
	"articles/internal/domain/models"
	"articles/internal/storage"
	"articles/pkg/jwt"
	"articles/pkg/logger/sl"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type Auth struct {
	Log          *slog.Logger
	UserSaver    UserSaver
	UserProvider UserProvider
	TokenTTl     time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (int64, error)
	MakeAdmin(ctx context.Context, userID int64) error
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

var (
	ErrUserExists         = errors.New("user does not exist")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func New(log *slog.Logger, userSaver UserSaver, userProvider UserProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		Log:          log,
		UserSaver:    userSaver,
		UserProvider: userProvider,
		TokenTTl:     tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email, password string) (string, error) {
	const op = "internal/services/blog/Login"
	log := a.Log.With(slog.String("op", op), slog.String("email", email))
	log.Info("attempting to login user")

	user, err := a.UserProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Warn("user not found", sl.Err(err))
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		log.Error("failed to get user", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err = bcrypt.CompareHashAndPassword(user.PassHas, []byte(password)); err != nil {
		log.Info("invalid credentials", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user successfully logged")

	token, err := jwt.New(user, a.TokenTTl)
	if err != nil {
		a.Log.Error("failed to generate token", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email, password string) (int64, error) {
	const op = "internal/services/blog/RegisterNewUser"
	log := a.Log.With(slog.String("op", op), slog.String("email", email))
	log.Info("registering new user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.UserSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Error("error", sl.Err(err))
		if errors.Is(err, storage.ErrUserExists) {
			log.With("user already exists", sl.Err(err))
			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		log.Error("failed to register user", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered")
	return id, nil
}

func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "internal/services/blog/IsAdmin"
	log := a.Log.With(slog.String("op", op), slog.Int64("user_id", userID))
	log.Info("checking if user is admin")

	isAdmin, err := a.UserProvider.IsAdmin(ctx, userID)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Warn("user not found", sl.Err(err))
			return false, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checked if user is admin", slog.Bool("is_admin", isAdmin))
	return isAdmin, nil
}

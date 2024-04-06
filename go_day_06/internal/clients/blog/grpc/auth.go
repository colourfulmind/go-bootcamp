package grpcclient

import (
	"articles/internal/ewrap"
	"articles/pkg/logger/sl"
	blog "articles/protos/gen/go/articles"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) RegisterNewUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "internal/clients/blog/RegisterNewUser"
		if r.Method == http.MethodGet {

			ts, err := template.ParseFiles("./static/register.html")
			if err != nil {
				c.Log.Error("cannot parse page", op, sl.Err(err))
				return
			}
			ts.Execute(w, "")
		} else if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				c.Log.Error("failed to get parameters", op, sl.Err(err))
				return
			}

			email := r.FormValue("u")
			password := r.FormValue("p")

			resp, err := c.Auth.RegisterNewUser(context.Background(), &blog.RegisterRequest{
				Email:    email,
				Password: password,
			})

			if err != nil {
				if errors.Is(err, ewrap.EmailIsRequired) || errors.Is(err, ewrap.PasswordIsRequired) {
					c.Log.Warn("email or password is required", op, sl.Err(err))
				} else if errors.Is(err, ewrap.UserAlreadyExists) {
					c.Log.Warn("user already exists", op, sl.Err(err))
				} else {
					c.Log.Warn("failed to register new user", err)
				}
			} else {
				c.Log.Info("user is registered", slog.Int64("user_id", resp.UserId))
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			c.Log.Error("method is not allowed", slog.String("method", r.Method))
		}
	}
}

func (c *Client) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "internal/clients/blog/Login"
		if limit.Allow() {
			if r.Method == http.MethodGet {

				ts, err := template.ParseFiles("./static/login.html")
				if err != nil {
					c.Log.Error("cannot parse page", op, sl.Err(err))
					return
				}
				ts.Execute(w, "")
			} else if r.Method == http.MethodPost {
				err := r.ParseForm()
				if err != nil {
					c.Log.Error("failed to get parameters", op, sl.Err(err))
					return
				}

				email := r.FormValue("u")
				password := r.FormValue("p")

				resp, err := c.Auth.Login(context.Background(), &blog.LoginRequest{
					Email:    email,
					Password: password,
				})

				if err != nil {
					if errors.Is(err, ewrap.EmailIsRequired) || errors.Is(err, ewrap.PasswordIsRequired) {
						c.Log.Warn("email or password is required", op, sl.Err(err))
					} else if errors.Is(err, ewrap.InvalidEmailOrPassword) {
						c.Log.Warn("invalid email or password", op, sl.Err(err))
					} else {
						c.Log.Error("failed to login user", err)
					}
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}

				c.Token = strings.TrimSuffix(strings.TrimPrefix(resp.String(), "token:\""), "\"")
				c.Log.Info("user successfully logged in", slog.String("token", resp.Token))
				http.Redirect(w, r, "/articles/my/?page=1", http.StatusSeeOther)

			} else {
				c.Log.Error("method is not allowed", slog.String("method", r.Method))
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
		} else {
			c.Log.Warn("access_denied", slog.Int("status", http.StatusTooManyRequests))
			http.Error(w, "access_denied", http.StatusTooManyRequests)
		}
	}
}

func ParseToken(token string, log *slog.Logger) (jwt.MapClaims, error) {
	const op = "internal/clients/blog/ParseToken"

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})

	log.Info("parsed", slog.Any("token", tokenParsed))
	if err != nil {
		log.Error("error parsing", op, slog.String("error", err.Error()))
		return jwt.MapClaims{}, err
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, errors.New("cannot get claims")
	}

	return claims, nil
}

func (c *Client) IsAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			const op = "internal/clients/blog/IsAdmin"

			id, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil || id == 0 {
				c.Log.Error("failed to get user id", op, sl.Err(err))
				return
			}

			resp, err := c.Auth.IsAdmin(context.Background(), &blog.IsAdminRequest{
				UserId: int64(id),
			})

			if err != nil {
				if errors.Is(err, ewrap.UserIDIsRequired) {
					c.Log.Warn("couldn't get user id")
					return
				}

				if errors.Is(err, ewrap.UserNotFound) {
					c.Log.Warn("user not found", op, sl.Err(err))
					return
				}

				c.Log.Error("failed to get user", err)
				return
			}

			c.Log.Info("checked if user is admin", slog.Bool("isAdmin", resp.IsAdmin))
			w.Write([]byte(resp.String()))
		} else {
			c.Log.Error("method is not allowed", slog.String("method", r.Method))
		}
	}
}

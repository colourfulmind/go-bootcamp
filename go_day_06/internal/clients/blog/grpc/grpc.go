package grpcclient

import (
	"articles/internal/config"
	blog "articles/protos/gen/go/articles"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	Auth     blog.AuthClient
	Articles blog.ArticlesClient
	Log      *slog.Logger
	Router   *mux.Router
	Config   *config.Config
	Token    string
}

var limit = rate.NewLimiter(rate.Limit(100), 1)

func New(cc *grpc.ClientConn, log *slog.Logger, config *config.Config) *Client {
	c := &Client{
		Auth:     blog.NewAuthClient(cc),
		Articles: blog.NewArticlesClient(cc),
		Log:      log,
		Router:   mux.NewRouter(),
		Config:   config,
	}
	c.ConfigureRouter()
	return c
}

func (c *Client) ConfigureRouter() {
	c.Router.HandleFunc("/articles/my/", c.ShowAllMyArticles())
	c.Router.HandleFunc("/articles/all/", c.ShowAllArticles())
	c.Router.HandleFunc("/article/my", c.ShowMyArticle())
	c.Router.HandleFunc("/article/post", c.CreateArticle())
	c.Router.HandleFunc("/login", c.Login())
	c.Router.HandleFunc("/register", c.RegisterNewUser())
	c.Router.HandleFunc("/access_denied", c.RegisterNewUser())
}

func (c *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Router.ServeHTTP(w, r)
}

func (c *Client) Start() error {
	addr := c.Config.GRPC.Host + ":" + strconv.Itoa(c.Config.GRPC.Port)
	c.Log.Info("starting client", slog.String("addr", addr))

	return http.ListenAndServe(addr, c)
}

func NewConnection(ctx context.Context, log *slog.Logger, addr string, retriesCount int, timeout time.Duration) (*grpc.ClientConn, error) {
	const op = "internal/clients/blog/NewConnection"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}

	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return cc, nil
}

func InterceptorLogger(log *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, level grpclog.Level, msg string, fields ...any) {
		log.Log(ctx, slog.Level(level), msg, fields...)
	})
}

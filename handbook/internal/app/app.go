package app

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/mayamika/mai-lev/handbook/internal/handlers"
	"github.com/mayamika/mai-lev/handbook/internal/postgres"
	"github.com/mayamika/mai-lev/handbook/internal/repository"
	"github.com/mayamika/mai-lev/handbook/internal/service"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type HttpConfig struct {
	Addr string
}

type Config struct {
	fx.Out

	HTTP     HttpConfig
	Postgres postgres.Config
}

func New(config Config) *fx.App {
	return fx.New(
		fx.Supply(config),

		fx.Provide(newLogger),
		fx.Provide(postgres.NewPool),

		fx.Provide(repository.NewRepository),
		fx.Provide(func(repo repository.Repository) service.EntriesRepo {
			return repo
		}),

		fx.Provide(service.NewEntriesService),
		fx.Provide(func(es *service.EntriesService) handlers.EntriesService {
			return es
		}),
		fx.Provide(handlers.NewEntries),

		fx.Invoke(httpServer),
	)
}

func newLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Development = false

	return config.Build()
}

func loggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			logger := logger.Named(`http`)

			wrapped := middleware.NewWrapResponseWriter(rw, r.ProtoMajor)
			start := time.Now()
			defer func() {
				logger.Info("Served",
					zap.Int("status", wrapped.Status()),
					zap.Duration("took", time.Since(start)),
					zap.String("remote", r.RemoteAddr),
					zap.String("request", r.RequestURI),
					zap.String("method", r.Method),
				)
			}()

			h.ServeHTTP(rw, r)
		})
	}
}

func httpServer(
	config HttpConfig,

	lc fx.Lifecycle,
	shutdowner fx.Shutdowner,

	logger *zap.Logger,
	entries *handlers.Entries,
) {
	r := chi.NewRouter()
	r.Use(
		loggingMiddleware(logger),
	)
	r.Route(`/api`, func(api chi.Router) {
		api.Mount(`/entries`, entries.Routes())
	})

	corsOptions := cors.Options{
		AllowedMethods: []string{
			`GET`,
			`POST`,
			`PATCH`,
			`DELETE`,
		},
	}
	srv := http.Server{
		Addr:    config.Addr,
		Handler: cors.New(corsOptions).Handler(r),
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					logger.Error("server closed",
						zap.Error(err),
					)
					shutdowner.Shutdown()
				}
			}()
			return nil
		},
		OnStop: srv.Shutdown,
	})
}

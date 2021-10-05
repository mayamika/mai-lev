package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Config struct {
	Conn string
}

type Pool struct {
	*pgxpool.Pool
}

func NewPool(config Config, lc fx.Lifecycle, logger *zap.Logger) (*Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(config.Conn)
	if err != nil {
		return nil, fmt.Errorf("can't parse config: %w", err)
	}
	poolConfig.ConnConfig.Logger = zapadapter.NewLogger(
		logger.Named(`postgres`),
	)

	p := &Pool{}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			p.Pool, err = pgxpool.ConnectConfig(ctx, poolConfig)
			if err != nil {
				return fmt.Errorf("can't connect to database: %w", err)
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			done := make(chan struct{})
			go func() {
				defer close(done)
				p.Pool.Close()
			}()

			select {
			case <-ctx.Done():
				return err
			case <-done:
				return nil
			}
		},
	})

	return p, nil
}

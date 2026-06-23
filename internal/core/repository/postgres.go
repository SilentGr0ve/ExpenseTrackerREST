package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

func NewPool(ctx context.Context, config config.PostgresConfig) (*PostgresPool, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)

	pgxconfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse pgx config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxconfig)
	if err != nil {
		return nil, fmt.Errorf("create pgx pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgx pool ping: %w", err)
	}

	return &PostgresPool{
		Pool:      pool,
		opTimeout: config.Timeout,
	}, nil
}

func (p *PostgresPool) OpTimeout() time.Duration {
	return p.opTimeout
}

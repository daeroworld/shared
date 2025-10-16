package database

import (
	"context"
	"time"

	"github.com/daeroworld/whiscript-shared/configuration"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresqlWrapper struct {
	Pool *pgxpool.Pool
}

func ConnectPostgresqlDatabase(database *configuration.Database) *PostgresqlWrapper {
	config, err := pgxpool.ParseConfig(database.Uri)
	if err != nil {
		return nil
	}
	config.MaxConns = 20
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.HealthCheckPeriod = 30 * time.Second
	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, config)
	wrapper := &PostgresqlWrapper{
		Pool: pool,
	}
	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctxPing); err != nil {
		pool.Close()
		return nil
	}

	return wrapper
}

func (p *PostgresqlWrapper) CloseDatabbase() {
	p.Pool.Close()
}

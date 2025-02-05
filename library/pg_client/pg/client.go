package pg

import (
  "context"

  db "github.com/Genvekt/kudos-vault/library/pg_client"

  "github.com/jackc/pgx/v4/pgxpool"
  "github.com/pkg/errors"
)

type pgClient struct {
  masterDBC db.DB
}

// New creates general provider of db clients
func New(ctx context.Context, dsn string) (db.Client, error) {
  dbc, err := pgxpool.Connect(ctx, dsn)
  if err != nil {
    return nil, errors.Errorf("failed to connect to db: %v", err)
  }

  return &pgClient{
    masterDBC: NewDB(dbc),
  }, nil
}

// DB returns one db client
func (c *pgClient) DB() db.DB {

  return c.masterDBC
}

// Close closes all db clients
func (c *pgClient) Close() error {
  if c.masterDBC != nil {
    c.masterDBC.Close()
  }

  return nil
}

package env

import (
  "fmt"
  "os"

  "github.com/Genvekt/kudos-vault/service/auth/internal/config"
)

var _ config.PostgresConfig = (*postgresConfigEnv)(nil)

const (
  dbNameEnv     = "PG_DATABASE_NAME"
  dbUserEnv     = "PG_USER"
  dbPasswordEnv = "PG_PASSWORD"
  dbHostEnv     = "PG_HOST"
  dbPortEnv     = "PG_PORT"
  dbDSNTemplate = "host=%s port=%s dbname=%s user=%s password=%s sslmode=disable"
)

type postgresConfigEnv struct {
  dbName     string
  dbUser     string
  dbPassword string
  dbHost     string
  dbPort     string
}

// NewPostgresConfigEnv retrieves new postgresConfigEnv instance
func NewPostgresConfigEnv() (*postgresConfigEnv, error) {
  dbName := os.Getenv(dbNameEnv)
  if dbName == "" {
    return nil, fmt.Errorf("environment variable %s not set", dbNameEnv)
  }

  dbUser := os.Getenv(dbUserEnv)
  if dbUser == "" {
    return nil, fmt.Errorf("environment variable %s not set", dbUserEnv)
  }

  dbPassword := os.Getenv(dbPasswordEnv)
  if dbPassword == "" {
    return nil, fmt.Errorf("environment variable %s not set", dbPasswordEnv)
  }

  dbHost := os.Getenv(dbHostEnv)
  if dbHost == "" {
    return nil, fmt.Errorf("environment variable %s not set", dbHostEnv)
  }

  dbPort := os.Getenv(dbPortEnv)
  if dbPort == "" {
    return nil, fmt.Errorf("environment variable %s not set", dbPortEnv)
  }

  return &postgresConfigEnv{
    dbName:     dbName,
    dbUser:     dbUser,
    dbPassword: dbPassword,
    dbHost:     dbHost,
    dbPort:     dbPort,
  }, nil
}

func (p *postgresConfigEnv) DSN() string {
  return fmt.Sprintf(dbDSNTemplate, p.dbHost, p.dbPort, p.dbName, p.dbUser, p.dbPassword)
}

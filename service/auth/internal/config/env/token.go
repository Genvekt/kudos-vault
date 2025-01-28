package env

import (
  "context"
  "fmt"
  "os"
  "time"

  "github.com/Genvekt/kudos-vault/service/auth/internal/config"
)

const (
  accessTtlEnv    = "ACCESS_TOKEN_TTL"
  accessSecretEnv = "ACCESS_TOKEN_SECRET"

  refreshTtlEnv    = "REFRESH_TOKEN_TTL"
  refreshSecretEnv = "REFRESH_TOKEN_SECRET"

  defaultAccessTTLStr  = "15m"
  defaultRefreshTTLStr = "24h"
)

var _ config.TokenConfig = (*accessTokenEnvConfig)(nil)
var _ config.TokenConfig = (*refreshTokenEnvConfig)(nil)

type refreshTokenEnvConfig struct {
  ttl    time.Duration
  secret []byte
}

func NewRefreshTokenEnvConfig(ctx context.Context) (*refreshTokenEnvConfig, error) {
  ttlStr := os.Getenv(refreshTtlEnv)
  if ttlStr == "" {
    ttlStr = defaultRefreshTTLStr
  }

  ttl, err := time.ParseDuration(ttlStr)
  if err != nil {
    return nil, fmt.Errorf("invalid value for %s: %w", refreshTtlEnv, err)
  }

  secret := os.Getenv(refreshSecretEnv)
  if secret == "" {
    return nil, fmt.Errorf("no value for environment variable %s provided", refreshSecretEnv)
  }

  return &refreshTokenEnvConfig{
    ttl:    ttl,
    secret: []byte(secret),
  }, nil
}

func (c *refreshTokenEnvConfig) TTL() time.Duration {
  return c.ttl
}

func (c *refreshTokenEnvConfig) Secret() []byte {
  return c.secret
}

type accessTokenEnvConfig struct {
  ttl    time.Duration
  secret []byte
}

func NewAccessTokenEnvConfig(ctx context.Context) (*accessTokenEnvConfig, error) {
  ttlStr := os.Getenv(accessTtlEnv)
  if ttlStr == "" {
    ttlStr = defaultAccessTTLStr
  }

  ttl, err := time.ParseDuration(ttlStr)
  if err != nil {
    return nil, fmt.Errorf("invalid value for %s: %w", accessTtlEnv, err)
  }

  secret := os.Getenv(accessSecretEnv)
  if secret == "" {
    return nil, fmt.Errorf("no value for environment variable %s provided", accessSecretEnv)
  }

  return &accessTokenEnvConfig{
    ttl:    ttl,
    secret: []byte(secret),
  }, nil
}

func (c *accessTokenEnvConfig) TTL() time.Duration {
  return c.ttl
}

func (c *accessTokenEnvConfig) Secret() []byte {
  return c.secret
}

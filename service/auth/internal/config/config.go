package config

import "time"

type TokenConfig interface {
  Secret() []byte
  TTL() time.Duration
}

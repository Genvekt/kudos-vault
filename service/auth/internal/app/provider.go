package app

import (
  "context"
  "log"

  db "github.com/Genvekt/kudos-vault/library/pg_client"
  "github.com/Genvekt/kudos-vault/library/pg_client/pg"
  authApi "github.com/Genvekt/kudos-vault/service/auth/internal/api/auth"
  userApi "github.com/Genvekt/kudos-vault/service/auth/internal/api/user"
  "github.com/Genvekt/kudos-vault/service/auth/internal/config"
  envConf "github.com/Genvekt/kudos-vault/service/auth/internal/config/env"
  "github.com/Genvekt/kudos-vault/service/auth/internal/repository"
  "github.com/Genvekt/kudos-vault/service/auth/internal/service"
  "github.com/Genvekt/kudos-vault/service/auth/internal/utils"
  "github.com/Genvekt/kudos-vault/service/auth/internal/utils/token"

  "github.com/Genvekt/kudos-vault/service/auth/internal/utils/hash"

  userRepo "github.com/Genvekt/kudos-vault/service/auth/internal/repository/user/postgres"
  authServ "github.com/Genvekt/kudos-vault/service/auth/internal/service/auth"
  userServ "github.com/Genvekt/kudos-vault/service/auth/internal/service/user"
)

type Provider struct {
  authImpl *authApi.Implementation
  userImpl *userApi.Implementation

  userService service.UserService
  authService service.AuthService

  userRepository repository.UserRepository

  pgClient db.Client

  hasher                utils.Hasher
  accessTockenProvider  utils.TokenProvider
  refreshTockenProvider utils.TokenProvider

  accessTokenConfig  config.TokenConfig
  refreshTokenConfig config.TokenConfig
  pgConfig           config.PostgresConfig
}

func newProvider() *Provider {
  return &Provider{}
}

func (p *Provider) AuthImpl(ctx context.Context) *authApi.Implementation {
  if p.authImpl == nil {
    p.authImpl = authApi.NewImplementation(p.AuthService(ctx))
  }

  return p.authImpl
}

func (p *Provider) UserImpl(ctx context.Context) *userApi.Implementation {
  if p.userImpl == nil {
    p.userImpl = userApi.NewImplementation(p.UserService(ctx))
  }

  return p.userImpl
}

func (p *Provider) UserService(ctx context.Context) service.UserService {
  if p.userService == nil {
    p.userService = userServ.NewUserService(p.UserRepo(ctx), p.Hasher(ctx))
  }

  return p.userService
}

func (p *Provider) AuthService(ctx context.Context) service.AuthService {
  if p.authService == nil {
    p.authService = authServ.NewAuthService(
      ctx,
      p.UserService(ctx),
      p.RefreshTokenProvider(ctx),
      p.AccessTokenProvider(ctx),
      p.Hasher(ctx),
    )
  }

  return p.authService
}

func (p *Provider) UserRepo(ctx context.Context) repository.UserRepository {
  if p.userRepository == nil {
    p.userRepository = userRepo.NewUserPgRepository(p.PgClient(ctx))
  }

  return p.userRepository
}

func (p *Provider) Hasher(ctx context.Context) utils.Hasher {
  if p.hasher == nil {
    p.hasher = hash.NewHasher()
  }

  return p.hasher
}

func (p *Provider) RefreshTokenProvider(ctx context.Context) utils.TokenProvider {
  if p.refreshTockenProvider == nil {
    p.refreshTockenProvider = token.NewTokenProvider(ctx, p.RefreshTokenConfig(ctx))
  }

  return p.refreshTockenProvider
}

func (p *Provider) AccessTokenProvider(ctx context.Context) utils.TokenProvider {
  if p.accessTockenProvider == nil {
    p.accessTockenProvider = token.NewTokenProvider(ctx, p.AccessTokenConfig(ctx))
  }

  return p.accessTockenProvider
}

func (p *Provider) RefreshTokenConfig(ctx context.Context) config.TokenConfig {
  if p.refreshTokenConfig == nil {
    conf, err := envConf.NewRefreshTokenEnvConfig(ctx)
    if err != nil {
      log.Fatal(err)
    }

    p.refreshTokenConfig = conf
  }

  return p.refreshTokenConfig
}

func (p *Provider) AccessTokenConfig(ctx context.Context) config.TokenConfig {
  if p.accessTokenConfig == nil {
    conf, err := envConf.NewAccessTokenEnvConfig(ctx)
    if err != nil {
      log.Fatal(err)
    }

    p.accessTokenConfig = conf
  }

  return p.accessTokenConfig
}

func (p *Provider) PgClient(ctx context.Context) db.Client {
  if p.pgClient == nil {

    pgClient, err := pg.New(ctx, p.PGConfig().DSN())
    if err != nil {
      log.Fatalf("failed to connect to postgres: %v", err)
    }

    if err := pgClient.DB().Ping(ctx); err != nil {
      log.Fatalf("failed to ping postgres %v", err)
    }

    p.pgClient = pgClient
  }

  return p.pgClient
}

func (p *Provider) PGConfig() config.PostgresConfig {
  if p.pgConfig == nil {
    conf, err := envConf.NewPostgresConfigEnv()
    if err != nil {
      log.Fatal(err)
    }

    p.pgConfig = conf
  }

  return p.pgConfig
}

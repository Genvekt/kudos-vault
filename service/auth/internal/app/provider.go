package app

import (
  "context"

  authApi "github.com/Genvekt/kudos-vault/service/auth/internal/api/auth"
  userApi "github.com/Genvekt/kudos-vault/service/auth/internal/api/user"
)

type Provider struct {
  authImpl *authApi.Implementation
  userImpl *userApi.Implementation
}

func newProvider() *Provider {
  return &Provider{}
}

func (p *Provider) AuthImpl(ctx context.Context) *authApi.Implementation {
  if p.authImpl == nil {
    p.authImpl = authApi.NewImplementation()
  }

  return p.authImpl
}

func (p *Provider) UserImpl(ctx context.Context) *userApi.Implementation {
  if p.userImpl == nil {
    p.userImpl = userApi.NewImplementation()
  }

  return p.userImpl
}

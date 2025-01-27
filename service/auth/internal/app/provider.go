package app

import (
  "context"

  authApi "github.com/Genvekt/kudos-vault/service/auth/internal/api/auth"
  userApi "github.com/Genvekt/kudos-vault/service/auth/internal/api/user"
  "github.com/Genvekt/kudos-vault/service/auth/internal/repository"
  "github.com/Genvekt/kudos-vault/service/auth/internal/service"

  userRepo "github.com/Genvekt/kudos-vault/service/auth/internal/repository/in_memo/user"
  userServ "github.com/Genvekt/kudos-vault/service/auth/internal/service/user"
)

type Provider struct {
  authImpl *authApi.Implementation
  userImpl *userApi.Implementation

  userService    service.UserService
  userRepository repository.UserRepository
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
    p.userImpl = userApi.NewImplementation(p.UserService(ctx))
  }

  return p.userImpl
}

func (p *Provider) UserService(ctx context.Context) service.UserService {
  if p.userService == nil {
    p.userService = userServ.NewService(p.UserRepo(ctx))
  }

  return p.userService
}

func (p *Provider) UserRepo(ctx context.Context) repository.UserRepository {
  if p.userRepository == nil {
    p.userRepository = userRepo.NewInMemoryRepository()
  }

  return p.userRepository
}

package service

import (
  "context"

  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
)

type UserService interface {
  Create(ctx context.Context, user *model.User, password string) (string, error)
  GetByID(ctx context.Context, id string) (*model.User, error)
  GetByEmail(ctx context.Context, email string) (*model.User, error)
  GetList(ctx context.Context, filters *model.UserListFilters) ([]*model.User, error)
}

type AuthService interface {
  Login(ctx context.Context, username string, password string) (string, string, error)
  GetRefreshToken(ctx context.Context, oldToken string) (string, error)
  GetAccessToken(ctx context.Context, oldToken string) (string, error)
}

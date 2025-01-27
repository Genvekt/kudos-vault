package service

import (
  "context"

  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
)

type UserService interface {
  Create(ctx context.Context, user *model.User, password string) (string, error)
  Get(ctx context.Context, id string) (*model.User, error)
  GetList(ctx context.Context, filters *model.UserListFilters) ([]*model.User, error)
}

package repository

import (
  "context"

  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
)

type UserRepository interface {
  Create(ctx context.Context, user *model.User) error
  GetByID(ctx context.Context, id string) (*model.User, error)
  GetByEmail(ctx context.Context, email string) (*model.User, error)
  GetList(ctx context.Context, filters *model.UserListFilters) ([]*model.User, error)
}

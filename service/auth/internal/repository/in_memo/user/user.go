package user

import (
  "context"
  "errors"
  "sync"

  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
  "github.com/Genvekt/kudos-vault/service/auth/internal/repository"
)

var _ repository.UserRepository = (*inMemoryUserRepository)(nil)

type inMemoryUserRepository struct {
  mu    sync.RWMutex
  users map[string]*model.User
}

func NewInMemoryRepository() *inMemoryUserRepository {
  return &inMemoryUserRepository{
    users: make(map[string]*model.User),
  }
}

func (r *inMemoryUserRepository) Create(ctx context.Context, user *model.User) error {
  r.mu.Lock()
  defer r.mu.Unlock()

  r.users[user.ID] = user

  return nil
}

func (r *inMemoryUserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
  r.mu.RLock()
  defer r.mu.RUnlock()
  user, exists := r.users[id]
  if !exists {
    return nil, errors.New("user not found")
  }
  return user, nil
}

func (r *inMemoryUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
  r.mu.RLock()
  defer r.mu.RUnlock()

  for _, user := range r.users {
    if user.Email == email {
      return user, nil
    }
  }

  return nil, errors.New("user not found")
}

func (r *inMemoryUserRepository) GetList(ctx context.Context, filters *model.UserListFilters) ([]*model.User, error) {
  r.mu.RLock()
  defer r.mu.RUnlock()
  var users []*model.User

  for _, user := range r.users {
    if filters != nil && len(filters.IDs) > 0 {
      for _, id := range filters.IDs {
        if user.ID == id {
          users = append(users, user)
        }
      }
    } else {
      users = append(users, user)
    }
  }

  return users, nil
}

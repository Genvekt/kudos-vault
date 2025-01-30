package user

import (
  "context"
  "fmt"
  "github.com/google/uuid"

  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
  "github.com/Genvekt/kudos-vault/service/auth/internal/repository"
  "github.com/Genvekt/kudos-vault/service/auth/internal/service"
  "github.com/Genvekt/kudos-vault/service/auth/internal/utils"
)

var _ service.UserService = (*userService)(nil)

type userService struct {
  repo   repository.UserRepository
  hasher utils.Hasher
}

func NewUserService(
  repo repository.UserRepository,
  hasher utils.Hasher,
) *userService {
  return &userService{repo: repo, hasher: hasher}
}

func (s *userService) Create(ctx context.Context, user *model.User, password string) (string, error) {
  id, err := generateID()
  if err != nil {
    return "", fmt.Errorf("failed to create uuid: %w", err)
  }

  user.ID = id

  passwordHash, err := s.hasher.HashPassword(ctx, password)
  if err != nil {
    return "", fmt.Errorf("failed to hash password: %w", err)
  }

  user.PasswordHash = passwordHash

  err = s.repo.Create(ctx, user)
  if err != nil {
    return "", err
  }

  return user.ID, nil
}

func (s *userService) GetByID(ctx context.Context, id string) (*model.User, error) {
  user, err := s.repo.GetByID(ctx, id)
  if err != nil {
    return nil, err
  }

  return user, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
  user, err := s.repo.GetByEmail(ctx, email)
  if err != nil {
    return nil, err
  }

  return user, nil
}

func (s *userService) GetList(ctx context.Context, filters *model.UserListFilters) ([]*model.User, error) {
  users, err := s.repo.GetList(ctx, filters)
  if err != nil {
    return nil, err
  }

  return users, nil
}

func generateID() (string, error) {
  rawUuid, err := uuid.NewUUID()
  if err != nil {
    return "", err
  }

  return rawUuid.String(), nil
}

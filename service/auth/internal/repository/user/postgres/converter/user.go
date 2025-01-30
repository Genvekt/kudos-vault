package converter

import (
  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
  repoModel "github.com/Genvekt/kudos-vault/service/auth/internal/repository/user/postgres/model"
)

// UserToRepositoryModel convert from internal service model to repository model
func UserToRepositoryModel(user *model.User) *repoModel.User {
  return &repoModel.User{
    ID:           user.ID,
    FirstName:    user.FirstName,
    LastName:     user.LastName,
    Email:        user.Email,
    PasswordHash: user.PasswordHash,
    Role:         user.Role,
    CreatedAt:    user.CreatedAt,
    UpdatedAt:    user.UpdatedAt,
    Status:       user.Status,
  }
}

// UserToServiceModel convert from repository model to internal service model
func UserToServiceModel(userRepo *repoModel.User) *model.User {
  return &model.User{
    ID:           userRepo.ID,
    FirstName:    userRepo.FirstName,
    LastName:     userRepo.LastName,
    Email:        userRepo.Email,
    PasswordHash: userRepo.PasswordHash,
    Role:         userRepo.Role,
    CreatedAt:    userRepo.CreatedAt,
    UpdatedAt:    userRepo.UpdatedAt,
    Status:       userRepo.Status,
  }
}

func UsersToServiceModel(users []*repoModel.User) []*model.User {
  result := make([]*model.User, len(users))
  for i, user := range users {
    result[i] = UserToServiceModel(user)
  }
  return result
}

func ListFiltersToRepoListFilters(filters *model.UserListFilters) *repoModel.ListFilters {
  return &repoModel.ListFilters{
    IDs: filters.IDs,
  }
}

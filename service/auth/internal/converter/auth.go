package converter

import (
  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
)

func ClaimsToUser(claims *model.UserClaims) *model.User {
  return &model.User{
    ID:        claims.ID,
    FirstName: claims.FirstName,
    LastName:  claims.LastName,
    Role:      claims.Role,
    Status:    claims.Status,
  }
}

func UserToClaims(user *model.User) *model.UserClaims {
  return &model.UserClaims{
    ID:        user.ID,
    FirstName: user.FirstName,
    LastName:  user.LastName,
    Role:      user.Role,
    Status:    user.Status,
  }
}

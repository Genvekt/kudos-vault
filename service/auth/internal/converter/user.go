package converter

import (
  "google.golang.org/protobuf/types/known/timestamppb"

  userApi "github.com/Genvekt/kudos-vault/library/api/user/v1"
  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
)

// FromUserToProtoUser convert internal user model to protobuf user
func FromUserToProtoUser(u *model.User) *userApi.User {
  if u == nil {
    return nil
  }

  return &userApi.User{
    Id: u.ID,
    Data: &userApi.UserData{
      FirstName: u.FirstName,
      LastName:  u.LastName,
      Email:     u.Email,
      Role:      userApi.UserRole(u.Role),
    },
    CreatedAt: timestamppb.New(u.CreatedAt),
    UpdatedAt: timestamppb.New(u.UpdatedAt),
    Status:    userApi.UserStatus(u.Status),
  }
}

func FromProtoListFiltersToListFilters(protoFilters *userApi.ListFilters) *model.UserListFilters {
  return &model.UserListFilters{
    IDs: protoFilters.GetIds(),
  }
}

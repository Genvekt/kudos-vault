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

  userRoleID := userApi.UserRole_value[u.Role]
  userStatusID := userApi.UserStatus_value[u.Status]

  return &userApi.User{
    Id: u.ID,
    Data: &userApi.UserData{
      FirstName: u.FirstName,
      LastName:  u.LastName,
      Email:     u.Email,
      Role:      userApi.UserRole(userRoleID),
    },
    CreatedAt: timestamppb.New(u.CreatedAt),
    UpdatedAt: timestamppb.New(u.UpdatedAt),
    Status:    userApi.UserStatus(userStatusID),
  }
}

func FromProtoListFiltersToListFilters(protoFilters *userApi.ListFilters) *model.UserListFilters {
  return &model.UserListFilters{
    IDs: protoFilters.GetIds(),
  }
}

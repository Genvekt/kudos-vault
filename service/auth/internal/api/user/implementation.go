package user

import (
  userApi "github.com/Genvekt/kudos-vault/library/api/user/v1"
)

type Implementation struct {
  userApi.UnimplementedUserV1Server
}

func NewImplementation() *Implementation {
  return &Implementation{}
}

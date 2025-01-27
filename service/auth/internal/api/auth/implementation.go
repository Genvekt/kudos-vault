package auth

import (
  authApi "github.com/Genvekt/kudos-vault/library/api/auth/v1"
)

type Implementation struct {
  authApi.UnimplementedAuthV1Server
}

func NewImplementation() *Implementation {
  return &Implementation{}
}

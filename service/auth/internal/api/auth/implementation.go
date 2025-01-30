package transport

import (
  "context"

  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"

  authApi "github.com/Genvekt/kudos-vault/library/api/auth/v1"
  "github.com/Genvekt/kudos-vault/service/auth/internal/service"
)

type Implementation struct {
  authApi.UnimplementedAuthV1Server
  service service.AuthService
}

func NewImplementation(svc service.AuthService) *Implementation {
  return &Implementation{service: svc}
}

func (s *Implementation) Login(ctx context.Context, req *authApi.LoginRequest) (*authApi.LoginResponse, error) {
  refreshToken, accessToken, err := s.service.Login(ctx, req.Username, req.Password)
  if err != nil {
    return nil, status.Error(codes.Unauthenticated, err.Error())
  }

  return &authApi.LoginResponse{
    RefreshToken: refreshToken,
    AccessToken:  accessToken,
  }, nil
}

func (s *Implementation) GetRefreshToken(ctx context.Context, req *authApi.GetRefreshTokenRequest) (*authApi.GetRefreshTokenResponse, error) {
  refreshToken, err := s.service.GetRefreshToken(ctx, req.OldRefreshToken)
  if err != nil {
    return nil, status.Error(codes.Unauthenticated, err.Error())
  }

  return &authApi.GetRefreshTokenResponse{
    RefreshToken: refreshToken,
  }, nil
}

func (s *Implementation) GetAccessToken(ctx context.Context, req *authApi.GetAccessTokenRequest) (*authApi.GetAccessTokenResponse, error) {
  accessToken, err := s.service.GetAccessToken(ctx, req.RefreshToken)
  if err != nil {
    return nil, status.Error(codes.Unauthenticated, err.Error())
  }

  return &authApi.GetAccessTokenResponse{
    AccessToken: accessToken,
  }, nil
}

package service

import (
  "context"
  "errors"

  "github.com/Genvekt/kudos-vault/service/auth/internal/converter"
  "github.com/Genvekt/kudos-vault/service/auth/internal/service"
  "github.com/Genvekt/kudos-vault/service/auth/internal/utils"
)

var _ service.AuthService = (*authService)(nil)

type authService struct {
  userService          service.UserService
  refreshTokenProvider utils.TokenProvider
  accessTokenProvider  utils.TokenProvider
  hasher               utils.Hasher
}

func NewAuthService(
  ctx context.Context,
  userService service.UserService,
  refreshTokenProvider utils.TokenProvider,
  accessTokenProvider utils.TokenProvider,
  hasher utils.Hasher,
) *authService {
  return &authService{
    userService:          userService,
    refreshTokenProvider: refreshTokenProvider,
    accessTokenProvider:  accessTokenProvider,
    hasher:               hasher,
  }
}

func (s *authService) Login(ctx context.Context, email, password string) (string, string, error) {
  // Fetch user by username
  user, err := s.userService.GetByEmail(ctx, email)
  if err != nil {
    return "", "", errors.New("invalid credentials")
  }

  // Validate password (pseudo-code)
  if !s.hasher.CheckPasswordHash(ctx, password, user.PasswordHash) {
    return "", "", errors.New("invalid credentials")
  }

  // Generate tokens
  refreshToken, err := s.refreshTokenProvider.Generate(ctx, user)
  if err != nil {
    return "", "", err
  }

  accessToken, err := s.accessTokenProvider.Generate(ctx, user)
  if err != nil {
    return "", "", err
  }

  return refreshToken, accessToken, nil
}

func (s *authService) GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
  // Validate old refresh token
  claims, err := s.refreshTokenProvider.Verify(ctx, oldRefreshToken)
  if err != nil {
    return "", err
  }

  // Generate new refresh token
  newRefreshToken, err := s.refreshTokenProvider.Generate(ctx, converter.ClaimsToUser(claims))
  if err != nil {
    return "", err
  }

  return newRefreshToken, nil
}

func (s *authService) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
  // Validate refresh token
  claims, err := s.refreshTokenProvider.Verify(ctx, refreshToken)
  if err != nil {
    return "", err
  }

  // Generate access token
  accessToken, err := s.accessTokenProvider.Generate(ctx, converter.ClaimsToUser(claims))
  if err != nil {
    return "", err
  }

  return accessToken, nil
}

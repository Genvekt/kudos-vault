package model

import "github.com/golang-jwt/jwt"

// UserClaims is data for token
type UserClaims struct {
  jwt.StandardClaims
  ID        string `json:"id"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Role      string `json:"role"`
  Status    string `json:"status"`
}

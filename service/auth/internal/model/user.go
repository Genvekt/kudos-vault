package model

import "time"

type User struct {
  ID           string
  FirstName    string
  LastName     string
  Email        string
  PasswordHash string
  CreatedAt    time.Time
  UpdatedAt    time.Time
  Status       string
}

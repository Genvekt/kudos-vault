package model

import "time"

type User struct {
  ID           string
  FirstName    string
  LastName     string
  Email        string
  PasswordHash string
  Role         int
  CreatedAt    time.Time
  UpdatedAt    time.Time
  Status       int
}

type UserListFilters struct {
  IDs []string
}

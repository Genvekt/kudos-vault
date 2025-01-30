package model

import "time"

type User struct {
  ID           string    `db:"id"`
  FirstName    string    `db:"first_name"`
  LastName     string    `db:"last_name"`
  Email        string    `db:"email"`
  PasswordHash string    `db:"password_hash"`
  Role         string    `db:"role"`
  CreatedAt    time.Time `db:"created_at"`
  UpdatedAt    time.Time `db:"updated_at"`
  Status       string    `db:"status"`
}

type ListFilters struct {
  IDs []string
}

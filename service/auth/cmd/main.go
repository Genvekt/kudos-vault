package main

import (
  "context"
  "log"

  "github.com/Genvekt/kudos-vault/service/auth/internal/app"
)

func main() {
  ctx := context.Background()
  app, err := app.NewApp(ctx)
  if err != nil {
    log.Fatal(err)
  }

  err = app.Run(ctx)
  if err != nil {
    log.Fatal(err)
  }
}

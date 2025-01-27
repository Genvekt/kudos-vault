package app

import (
  "context"
  "fmt"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  "google.golang.org/grpc/reflection"
  "log"
  "net"
  "sync"

  authApi "github.com/Genvekt/kudos-vault/library/api/auth/v1"
  userApi "github.com/Genvekt/kudos-vault/library/api/user/v1"
)

type App struct {
  provider   *Provider
  grpcServer *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
  application := &App{}

  err := application.initDeps(ctx)
  if err != nil {
    return nil, err
  }

  return application, nil
}

func (a *App) initDeps(ctx context.Context) error {
  deps := []func(context.Context) error{
    a.initServiceProvider,
    a.initGRPCServer,
  }

  for _, dep := range deps {
    if err := dep(ctx); err != nil {
      return err
    }
  }

  return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
  a.provider = newProvider()

  return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
  // configure TLS if it is enabled
  creds := insecure.NewCredentials()

  a.grpcServer = grpc.NewServer(
    grpc.Creds(creds),
  )

  reflection.Register(a.grpcServer)

  userApi.RegisterUserV1Server(a.grpcServer, a.provider.UserImpl(ctx))
  authApi.RegisterAuthV1Server(a.grpcServer, a.provider.AuthImpl(ctx))

  return nil
}

func (a *App) runGRPCServer(_ context.Context) error {
  lis, err := net.Listen("tcp", "0.0.0.0:50051")
  if err != nil {

    return fmt.Errorf("failed to listen: %v", err)
  }

  err = a.grpcServer.Serve(lis)
  if err != nil {
    return err
  }

  return nil
}

func (a *App) Run(ctx context.Context) error {
  wg := sync.WaitGroup{}
  wg.Add(1)

  go func() {
    defer wg.Done()
    if err := a.runGRPCServer(ctx); err != nil {
      log.Fatalf("failed to run GRPC server: %v", err)
    }
  }()

  wg.Wait()
  return nil

}

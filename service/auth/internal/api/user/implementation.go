package user

import (
  "context"
  "time"

  userApi "github.com/Genvekt/kudos-vault/library/api/user/v1"
  "github.com/Genvekt/kudos-vault/service/auth/internal/converter"
  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
  "github.com/Genvekt/kudos-vault/service/auth/internal/service"
)

type Implementation struct {
  userApi.UnimplementedUserV1Server
  userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
  return &Implementation{
    userService: userService,
  }
}

// Create handles the Create RPC
func (i *Implementation) Create(ctx context.Context, req *userApi.CreateRequest) (*userApi.CreateResponse, error) {
  user := &model.User{
    FirstName: req.GetData().GetFirstName(),
    LastName:  req.GetData().GetLastName(),
    Email:     req.GetData().GetEmail(),
    Role:      req.GetData().GetRole().String(),
    Status:    userApi.UserStatus_STATUS_ACTIVE.String(), // Default status
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  id, err := i.userService.Create(ctx, user, req.GetPassword())
  if err != nil {
    return nil, err
  }

  return &userApi.CreateResponse{Id: id}, nil
}

// Get handles the Get RPC
func (i *Implementation) Get(ctx context.Context, req *userApi.GetRequest) (*userApi.GetResponse, error) {
  user, err := i.userService.GetByID(ctx, req.GetId())
  if err != nil {
    return nil, err
  }

  return &userApi.GetResponse{
    User: converter.FromUserToProtoUser(user),
  }, nil
}

// GetList handles the GetList RPC
func (i *Implementation) GetList(ctx context.Context, req *userApi.GetListRequest) (*userApi.GetListResponse, error) {
  filters := converter.FromProtoListFiltersToListFilters(req.GetFilters())
  users, err := i.userService.GetList(ctx, filters)
  if err != nil {
    return nil, err
  }

  protoUsers := make([]*userApi.User, len(users))
  for i, user := range users {
    protoUsers[i] = converter.FromUserToProtoUser(user)
  }

  return &userApi.GetListResponse{Users: protoUsers}, nil
}

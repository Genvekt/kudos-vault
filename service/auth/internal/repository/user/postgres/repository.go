package postgres

import (
  "context"
  "fmt"
  sq "github.com/Masterminds/squirrel"

  db "github.com/Genvekt/kudos-vault/library/pg_client"

  "github.com/Genvekt/kudos-vault/service/auth/internal/model"
  "github.com/Genvekt/kudos-vault/service/auth/internal/repository"
  "github.com/Genvekt/kudos-vault/service/auth/internal/repository/user/postgres/converter"
  repoModel "github.com/Genvekt/kudos-vault/service/auth/internal/repository/user/postgres/model"
)

const (
  usersTable         = "users"
  idColumn           = "id"
  firstNameColumn    = "first_name"
  lastNameColumn     = "last_name"
  emailColumn        = "email"
  passwordHashColumn = "password_hash"
  roleColumn         = "role"
  statusColumn       = "status"
)

var _ repository.UserRepository = (*userPgRepository)(nil)

type userPgRepository struct {
  db db.Client
}

func NewUserPgRepository(db db.Client) *userPgRepository {
  return &userPgRepository{
    db: db,
  }
}

func (r *userPgRepository) Create(ctx context.Context, user *model.User) error {
  dbUser := converter.UserToRepositoryModel(user)

  insertBuilder := sq.Insert(usersTable).
    PlaceholderFormat(sq.Dollar).
    Columns(
      idColumn,
      firstNameColumn,
      lastNameColumn,
      emailColumn,
      passwordHashColumn,
      roleColumn,
      statusColumn,
    ).
    Values(
      dbUser.ID,
      dbUser.FirstName,
      dbUser.LastName,
      dbUser.Email,
      dbUser.PasswordHash,
      dbUser.Role,
      dbUser.Status,
    )

  query, args, err := insertBuilder.ToSql()
  if err != nil {
    return fmt.Errorf("failed to build query: %v", err)
  }

  q := db.Query{
    Name:     "user_repository.Create",
    QueryRaw: query,
  }

  _, err = r.db.DB().ExecContext(ctx, q, args...)
  if err != nil {
    return err
  }

  return nil
}

func (r *userPgRepository) GetByID(ctx context.Context, userID string) (*model.User, error) {
  selectBuilder := sq.Select("*").
    From(usersTable).
    PlaceholderFormat(sq.Dollar).
    Where(sq.Eq{idColumn: userID}).
    Limit(1)

  query, args, err := selectBuilder.ToSql()
  if err != nil {
    return nil, fmt.Errorf("failed to build query: %v", err)
  }

  q := db.Query{
    Name:     "user_repository.GetById",
    QueryRaw: query,
  }

  user := &repoModel.User{}
  err = r.db.DB().ScanOneContext(ctx, user, q, args...)
  if err != nil {
    return nil, err
  }

  return converter.UserToServiceModel(user), nil
}

func (r *userPgRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
  selectBuilder := sq.Select("*").
    From(usersTable).
    PlaceholderFormat(sq.Dollar).
    Where(sq.Eq{emailColumn: email}).
    Limit(1)

  query, args, err := selectBuilder.ToSql()
  if err != nil {
    return nil, fmt.Errorf("failed to build query: %v", err)
  }

  q := db.Query{
    Name:     "user_repository.GetByEmail",
    QueryRaw: query,
  }

  user := &repoModel.User{}
  err = r.db.DB().ScanOneContext(ctx, user, q, args...)
  if err != nil {
    return nil, err
  }

  return converter.UserToServiceModel(user), nil
}

func (r *userPgRepository) GetList(ctx context.Context, filters *model.UserListFilters) ([]*model.User, error) {
  dbFilters := converter.ListFiltersToRepoListFilters(filters)

  selectBuilder := sq.Select("*").
    From(usersTable).
    PlaceholderFormat(sq.Dollar)

  if dbFilters != nil && len(dbFilters.IDs) > 0 {
    selectBuilder = selectBuilder.Where(sq.Eq{idColumn: dbFilters.IDs})
  }

  query, args, err := selectBuilder.ToSql()
  if err != nil {
    return nil, fmt.Errorf("failed to build query: %v", err)
  }

  q := db.Query{
    Name:     "user_repository.GetList",
    QueryRaw: query,
  }

  var users []*repoModel.User
  err = r.db.DB().ScanAllContext(ctx, &users, q, args...)
  if err != nil {
    return nil, err
  }

  return converter.UsersToServiceModel(users), nil
}

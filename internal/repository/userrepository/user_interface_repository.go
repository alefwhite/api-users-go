package userrepository

import (
	"context"
	"database/sql"
	"github.com/alefwhite/api-users-go/internal/database/sqlc"
	"github.com/alefwhite/api-users-go/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u *entity.UserEntity) error
	FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error)
	UpdateUser(ctx context.Context, u *entity.UserEntity) error
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) ([]entity.UserEntity, error)
	UpdatePassword(ctx context.Context, pass, id string) error
	GetUserPassword(ctx context.Context, id string) (string, error)
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewUserRepository(db *sql.DB, q *sqlc.Queries) UserRepository {
	return &repository{
		db,
		q,
	}
}

package userservice

import (
	"context"
	"github.com/alefwhite/api-users-go/internal/dto"
	"github.com/alefwhite/api-users-go/internal/handler/response"
	"github.com/alefwhite/api-users-go/internal/repository/userrepository"
)

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
	GetUserByID(ctx context.Context, id string) (*response.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) (*response.ManyUsersResponse, error)
	UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error
	Login(ctx context.Context, u dto.LoginDTO) (*response.UserAuthToken, error)
}

type service struct {
	repo userrepository.UserRepository
}

func NewUserService(repo userrepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

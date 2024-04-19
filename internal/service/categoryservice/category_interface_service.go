package categoryservice

import (
	"context"
	"github.com/alefwhite/api-users-go/internal/dto"
	"github.com/alefwhite/api-users-go/internal/repository/categoryrepository"
)

func NewCategoryService(repo categoryrepository.CategoryRepository) CategoryService {
	return &service{
		repo,
	}
}

type service struct {
	repo categoryrepository.CategoryRepository
}

type CategoryService interface {
	CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error
}

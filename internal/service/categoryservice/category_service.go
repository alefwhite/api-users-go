package categoryservice

import (
	"context"
	"errors"
	"github.com/alefwhite/api-users-go/internal/dto"
	"github.com/alefwhite/api-users-go/internal/entity"
	"github.com/google/uuid"
	"time"
)

func (s *service) CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error {
	categoryEntity := entity.CategoryEntity{
		ID:        uuid.New().String(),
		Title:     u.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.repo.CreateCategory(ctx, &categoryEntity)
	if err != nil {
		return errors.New("error to create category")
	}
	return nil
}

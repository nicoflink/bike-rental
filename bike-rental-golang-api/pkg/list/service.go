package list

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAllBikes(ctx context.Context, userID uuid.UUID) ([]Bike, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s Service) GetAllBikes(ctx context.Context, userID uuid.UUID) ([]Bike, error) {
	return s.repository.GetAllBikes(ctx, userID)
}

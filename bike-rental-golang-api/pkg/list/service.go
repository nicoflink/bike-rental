package list

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAllBikes(userID uuid.UUID) ([]Bike, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s Service) GetAllBikes(context context.Context) ([]Bike, error) {
	return s.repository.GetAllBikes(uuid.New())
}

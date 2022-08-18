package list

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	GetAllBikes(ctx context.Context, userID uuid.UUID) ([]Bike, error)
	GetListBikeByID(ctx context.Context, userID uuid.UUID, bikeID uuid.UUID) (Bike, error)
	UpdateBike(ctx context.Context, b Bike) (Bike, error)
}

type Service struct {
	repository Repository
}

// NewService returns a new bike service.
func NewService(r Repository) *Service {
	return &Service{repository: r}
}

// GetAllBikes for the current user.
func (s Service) GetAllBikes(ctx context.Context, userID uuid.UUID) ([]Bike, error) {
	return s.repository.GetAllBikes(ctx, userID)
}

// UpdateBikePosition updates the position of the bike.
func (s Service) UpdateBikePosition(ctx context.Context, userID uuid.UUID, locationUpdate BikeLocationUpdate) (Bike, error) {
	const prefix = "list.Service.UpdateBikePosition"

	movedBike, err := s.repository.GetListBikeByID(ctx, userID, locationUpdate.BikeID)
	if err != nil {
		return Bike{}, fmt.Errorf("%s: Unable to get bike: %w", prefix, err)
	}

	movedBike.Location = locationUpdate.Location

	updatedBike, err := s.repository.UpdateBike(ctx, movedBike)
	if err != nil {
		return Bike{}, fmt.Errorf("%s: Unable to update bike location: %w", prefix, err)
	}

	return updatedBike, nil
}

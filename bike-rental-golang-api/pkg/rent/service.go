package rent

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/persistence"
)

type Repository interface {
	GetBikeByID(ctx context.Context, userID uuid.UUID) (Bike, error)
	GetBikeByUserID(ctx context.Context, userID uuid.UUID) (Bike, error)
	GetRentByID(ctx context.Context, rentID uuid.UUID) (Rent, error)
	GetRentByStatusAndRenterID(_ context.Context, status Status, renter uuid.UUID) ([]Rent, error)
	CreateRentAndUpdateBike(ctx context.Context, r Rent, b Bike) (Rent, error)
	UpdateRentAndUpdateBike(ctx context.Context, r Rent, b Bike) (Rent, error)
}

type Service struct {
	repository Repository
}

// NewService returns a new rent service.
func NewService(r Repository) *Service {
	return &Service{repository: r}
}

// StartRent creates a new rent.
func (s Service) StartRent(ctx context.Context, startRequest StartRequest) (Rent, error) {
	const prefix = "rent.Service.StartRent"

	_, err := s.repository.GetBikeByUserID(ctx, startRequest.Renter)
	if err != nil && !errors.Is(err, persistence.ErrMissingResource) {
		return Rent{}, fmt.Errorf("%s: Unable to get bike: %w", prefix, err)
	}

	if err == nil {
		return Rent{}, fmt.Errorf("%s: User already rented a bike", prefix)
	}

	bike, err := s.repository.GetBikeByID(ctx, startRequest.Bike)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: Unable to get bike: %w", prefix, err)
	}

	r := NewRent(startRequest.Bike, startRequest.Renter, bike.Location)

	err = bike.startRent(r.Renter)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: cannot start rent: %w", prefix, err)
	}

	rCreated, err := s.repository.CreateRentAndUpdateBike(ctx, *r, bike)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: %w", prefix, err)
	}

	return rCreated, nil
}

// StopRent stops a started rent.
func (s Service) StopRent(ctx context.Context, stopRequest StopRequest) (Rent, error) {
	const prefix = "rent.Service.StopRent"

	ren, err := s.repository.GetRentByID(ctx, stopRequest.RentID)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: get Rent with ID %s : %w", prefix, stopRequest.RentID.String(), err)
	}

	bike, err := s.repository.GetBikeByID(ctx, ren.Bike)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: Unable to get bike: %w", prefix, err)
	}

	if err = ren.StopRent(stopRequest.UserID, bike.Location); err != nil {
		return Rent{}, fmt.Errorf("%s: cannot stop rent: %w", prefix, err)
	}

	if err = bike.stopRent(stopRequest.UserID); err != nil {
		return Rent{}, fmt.Errorf("%s: cannot stop rent: %w", prefix, err)
	}

	rUpdated, err := s.repository.UpdateRentAndUpdateBike(ctx, ren, bike)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.StopRent: update Rent with ID %s : %w", stopRequest.RentID.String(), err)
	}

	return rUpdated, nil
}

// GetStartedRents returns all rents of a user with status Started.
func (s Service) GetStartedRents(ctx context.Context, req GetOpenRentsRequest) ([]Rent, error) {
	return s.repository.GetRentByStatusAndRenterID(ctx, Started, req.UserID)
}

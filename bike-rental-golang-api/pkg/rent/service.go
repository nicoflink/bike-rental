package rent

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
	"github.com/nicoflink/bike-rental/pkg/persistence"
)

type Repository interface {
	GetBikeByID(ctx context.Context, userID uuid.UUID) (Bike, error)
	GetBikeByUserID(ctx context.Context, userID uuid.UUID) (Bike, error)
	GetRentByID(ctx context.Context, rentID uuid.UUID) (Rent, error)
	CreateRentAndUpdateBike(ctx context.Context, r Rent, b Bike) (Rent, error)
	UpdateRentAndUpdateBike(ctx context.Context, r Rent, b Bike) (Rent, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s Service) StartRent(ctx context.Context, request Request) (Rent, error) {
	const prefix = "rent.Service.StartRent"

	_, err := s.repository.GetBikeByUserID(ctx, request.Renter)
	if err != nil && !errors.Is(err, persistence.ErrMissingResource) {
		return Rent{}, fmt.Errorf("%s: Unable to get bike: %w", prefix, err)
	}

	if err == nil {
		return Rent{}, fmt.Errorf("%s: User already rented a bike", prefix)
	}

	bike, err := s.repository.GetBikeByID(ctx, request.Bike)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: Unable to get bike: %w", prefix, err)
	}

	if bike.RentedBy != nil {
		return Rent{}, fmt.Errorf("%s: Bike is already rented", prefix)
	}

	r := NewRent(request.Bike, request.Renter, bike.Location)
	bike.RentedBy = &request.Renter

	rCreated, err := s.repository.CreateRentAndUpdateBike(ctx, *r, bike)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: %w", prefix, err)
	}

	return rCreated, nil
}

func (s Service) StopRent(ctx context.Context, RentID uuid.UUID, endLocation geo.Coordinates) (Rent, error) {
	const prefix = "rent.Service.StopRent"

	ren, err := s.repository.GetRentByID(ctx, RentID)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: get Rent with ID %s : %w", prefix, RentID.String(), err)
	}

	bike, err := s.repository.GetBikeByID(ctx, ren.Bike)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: Unable to get bike: %w", prefix, err)
	}

	ren.StopRent(endLocation)
	bike.removeRenter()

	rUpdated, err := s.repository.UpdateRentAndUpdateBike(ctx, ren, bike)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.StopRent: update Rent with ID %s : %w", RentID.String(), err)
	}

	return rUpdated, nil
}

package rent

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

type Repository interface {
	GetBikeByID(context.Context, uuid.UUID) (Bike, error)
	GetRentByID(context.Context, uuid.UUID) (Rent, error)
	CreateRentAndUpdateBike(context.Context, Rent, Bike) (Rent, error)
	UpdateRentAndUpdateBike(context.Context, Rent, Bike) (Rent, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s Service) StartRent(ctx context.Context, userID uuid.UUID, bikeID uuid.UUID) (Rent, error) {
	const prefix = "rent.Service.StartRent"

	bike, err := s.repository.GetBikeByID(ctx, bikeID)
	if err != nil {
		return Rent{}, fmt.Errorf("%s: Unable to get bike: %w", prefix, err)
	}

	if bike.RentedBy != nil {
		return Rent{}, fmt.Errorf("%s: Bike is already rented", prefix)
	}

	r := NewRent(bikeID, userID, bike.Location)
	bike.RentedBy = &userID

	rCreated, err := s.repository.CreateRentAndUpdateBike(ctx, *r, bike)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.CreateRentAndUpdateBike: %w", err)
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

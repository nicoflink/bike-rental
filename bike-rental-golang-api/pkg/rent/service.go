package rent

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
)

type Repository interface {
	GetBikeByID(bikeID uuid.UUID) (Bike, error)
	UpdateBike(bike Bike) (Bike, error)
	GetRentByID(RentID uuid.UUID) (Rent, error)
	CreateRent(Rent Rent) (ren Rent, err error)
	UpdateRent(Rent Rent) (Rent, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s Service) StartRent(userID uuid.UUID, bikeID uuid.UUID) (Rent, error) {
	bike, err := s.repository.GetBikeByID(bikeID)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.StartRent: Unable to get bike: %w", err)
	}

	if bike.RentedBy != nil {
		return Rent{}, errors.New("rent.Service.StartRent: Bike is already rented")
	}

	r := NewRent(bikeID, userID, bike.Location)

	rCreated, err := s.repository.CreateRent(*r)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.StartRent: Unable to create rent: %w", err)
	}

	bike.RentedBy = &userID

	_, err = s.repository.UpdateBike(bike)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.StartRent: Unable to update bike: %w", err)
	}

	return rCreated, nil
}

func (s Service) StopRent(RentID uuid.UUID, endLocation geo.Coordinates) (Rent, error) {
	ren, err := s.repository.GetRentByID(RentID)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.StopRent: get Rent with ID %s : %w", RentID.String(), err)
	}

	ren.StopRent(endLocation)

	rUpdated, err := s.repository.UpdateRent(ren)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.StopRent: update Rent with ID %s : %w", RentID.String(), err)
	}

	bike, err := s.repository.GetBikeByID(ren.Bike)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.StartRent: Unable to get bike: %w", err)
	}

	bike.removeRenter()

	_, err = s.repository.UpdateBike(bike)
	if err != nil {
		return Rent{}, fmt.Errorf("rent.Service.StartRent: Unable to update bike: %w", err)
	}

	return rUpdated, nil
}

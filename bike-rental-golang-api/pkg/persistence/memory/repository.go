package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/list"
	"github.com/nicoflink/bike-rental/pkg/persistence"
	"github.com/nicoflink/bike-rental/pkg/rent"
)

type Repository struct {
	lock    sync.RWMutex
	bikes   map[uuid.UUID]Bike
	rentals map[uuid.UUID]rent.Rent
}

type RepositoryOption func(r *Repository)

func NewRepository(options ...RepositoryOption) *Repository {
	r := &Repository{
		bikes:   make(map[uuid.UUID]Bike),
		rentals: make(map[uuid.UUID]rent.Rent),
	}

	for _, f := range options {
		f(r)
	}

	return r
}

func WithSampleBikes(bikes []Bike) RepositoryOption {
	return func(r *Repository) {
		for _, b := range bikes {
			r.bikes[b.ID] = b
		}
	}
}

func (r *Repository) GetAllBikes(_ context.Context, userID uuid.UUID) ([]list.Bike, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	bikes := make([]list.Bike, 0, len(r.bikes))

	for _, b := range r.bikes {
		bikes = append(bikes, mapBikeToListBike(b, userID))
	}

	return bikes, nil
}

func (r *Repository) GetBikeByID(_ context.Context, bikeID uuid.UUID) (rent.Bike, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	b, ok := r.bikes[bikeID]
	if !ok {
		return rent.Bike{}, persistence.ErrMissingResource
	}

	return mapBikeToRentBike(b), nil
}

func (r *Repository) GetBikeByUserID(_ context.Context, userID uuid.UUID) (rent.Bike, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	for _, b := range r.bikes {
		if b.RentedByUser != nil && *b.RentedByUser == userID {
			return rent.Bike{}, nil
		}
	}

	return rent.Bike{}, persistence.ErrMissingResource
}

func (r *Repository) GetRentByID(_ context.Context, rentID uuid.UUID) (rent.Rent, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	ren, ok := r.rentals[rentID]
	if !ok {
		return rent.Rent{}, persistence.ErrMissingResource
	}

	return ren, nil
}

func (r *Repository) CreateRentAndUpdateBike(_ context.Context, ren rent.Rent, b rent.Bike) (rent.Rent, error) {
	const prefix = "memory.Repository.CreateRentAndUpdateBike"

	r.lock.Lock()
	defer r.lock.Unlock()

	r.rentals[ren.ID] = ren

	bike, ok := r.bikes[b.ID]
	if !ok {
		return rent.Rent{}, fmt.Errorf("%s : Unable to find bike to be updated", prefix)
	}

	bike.Location = b.Location
	bike.RentedByUser = b.RentedBy

	r.bikes[b.ID] = bike

	return ren, nil
}

func (r *Repository) UpdateRentAndUpdateBike(_ context.Context, ren rent.Rent, b rent.Bike) (rent.Rent, error) {
	const prefix = "memory.Repository.UpdateRentAndUpdateBike"

	r.lock.Lock()
	defer r.lock.Unlock()

	r.rentals[ren.ID] = ren

	bike, ok := r.bikes[b.ID]
	if !ok {
		return rent.Rent{}, fmt.Errorf("%s : Unable to find bike to be updated", prefix)
	}

	bike.Location = b.Location
	bike.RentedByUser = b.RentedBy

	r.bikes[b.ID] = bike

	return ren, nil
}

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

// Repository is the in memory storage for bikes and rentals.
type Repository struct {
	lock    sync.RWMutex
	bikes   map[uuid.UUID]Bike
	rentals map[uuid.UUID]rent.Rent
}

type RepositoryOption func(r *Repository)

// NewRepository initializes a new in memory storage with optional options.
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

// WithSampleBikes is the option to set initial sample bikes.
func WithSampleBikes(bikes []Bike) RepositoryOption {
	return func(r *Repository) {
		for _, b := range bikes {
			r.bikes[b.ID] = b
		}
	}
}

// WithSampleRents is the option to set initial sample rents.
func WithSampleRents(rents []rent.Rent) RepositoryOption {
	return func(r *Repository) {
		for _, ren := range rents {
			r.rentals[ren.ID] = ren
		}
	}
}

// GetAllBikes returns all bike mapped to list bike.
func (r *Repository) GetAllBikes(_ context.Context, userID uuid.UUID) ([]list.Bike, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	bikes := make([]list.Bike, 0, len(r.bikes))

	for _, b := range r.bikes {
		bikes = append(bikes, mapBikeToListBike(b, userID))
	}

	return bikes, nil
}

// GetBikeByID is a lookup for the bike with the provided bikeID and mapped to rent bike.
func (r *Repository) GetBikeByID(_ context.Context, bikeID uuid.UUID) (rent.Bike, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	b, ok := r.bikes[bikeID]
	if !ok {
		return rent.Bike{}, persistence.ErrMissingResource
	}

	return mapBikeToRentBike(b), nil
}

// GetListBikeByID is a lookup for the bike with the provided bikeID and mapped to list bike.
func (r *Repository) GetListBikeByID(_ context.Context, userID uuid.UUID, bikeID uuid.UUID) (list.Bike, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	b, ok := r.bikes[bikeID]
	if !ok {
		return list.Bike{}, persistence.ErrMissingResource
	}

	return mapBikeToListBike(b, userID), nil
}

// GetBikeByUserID is a lookup for the bike of a specific user.
func (r *Repository) GetBikeByUserID(_ context.Context, userID uuid.UUID) (rent.Bike, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	for _, b := range r.bikes {
		if b.RentedByUser != nil && *b.RentedByUser == userID {
			return mapBikeToRentBike(b), nil
		}
	}

	return rent.Bike{}, persistence.ErrMissingResource
}

// UpdateBike updates a list bike.
// Only Name and Location can be updated with this function.
func (r *Repository) UpdateBike(_ context.Context, b list.Bike) (list.Bike, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	bike, ok := r.bikes[b.ID]
	if !ok {
		return list.Bike{}, persistence.ErrMissingResource
	}

	bike.Name = b.Name
	bike.Location = b.Location

	r.bikes[b.ID] = bike

	return b, nil
}

// GetRentByID is a lookup for the rent with the provided ID.
func (r *Repository) GetRentByID(_ context.Context, rentID uuid.UUID) (rent.Rent, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	ren, ok := r.rentals[rentID]
	if !ok {
		return rent.Rent{}, persistence.ErrMissingResource
	}

	return ren, nil
}

// GetRentByStatusAndRenterID is a lookup for the rent with the provided status and renter.
func (r *Repository) GetRentByStatusAndRenterID(_ context.Context, status rent.Status, renter uuid.UUID) ([]rent.Rent, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	rents := make([]rent.Rent, 0, 1)

	for _, ren := range r.rentals {
		if ren.Status == status && ren.Renter == renter {
			rents = append(rents, ren)
		}
	}

	return rents, nil
}

// CreateRentAndUpdateBike creates a new rent and updates a bike in terms of rented and rentedByUser.
// Creation and Update are done in a single function as it should be handled as a transaction.
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

// UpdateRentAndUpdateBike updates a rent and a bike in terms of rented and rentedByUser.
// Both updated are done in a single function as it should be handled as a transaction.
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

package memory

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/geo"
	"github.com/nicoflink/bike-rental/pkg/list"
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

func (r *Repository) GetAllBikes(userID uuid.UUID) ([]list.Bike, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	bikes := make([]list.Bike, 0, len(r.bikes))

	for _, b := range r.bikes {
		var rentedByUser bool
		var rented bool

		if b.RentedByUser != nil {
			rented = true

			if *b.RentedByUser == userID {
				rentedByUser = true
			}
		}

		vBike := list.Bike{
			ID:   b.ID,
			Name: b.Name,
			Location: geo.Coordinates{
				Lat: b.Location.Lat,
				Lng: b.Location.Lng,
			},
			Rented:       rented,
			RentedByUser: rentedByUser,
		}

		bikes = append(bikes, vBike)
	}

	return bikes, nil
}

func (r *Repository) CreateRent(ren rent.Rent) (rent.Rent, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.rentals[ren.ID] = ren

	return ren, nil
}

func (r *Repository) UpdateRent(ren rent.Rent) (rent.Rent, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.rentals[ren.ID] = ren

	return ren, nil
}

func (r *Repository) getBikeByID(bikeID uuid.UUID) (Bike, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	b, ok := r.bikes[bikeID]
	if !ok {
		return Bike{}, errors.New("missing ressource")
	}

	return b, nil
}

func (r *Repository) updateBike(bike Bike) (Bike, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.bikes[bike.ID] = bike

	return bike, nil
}

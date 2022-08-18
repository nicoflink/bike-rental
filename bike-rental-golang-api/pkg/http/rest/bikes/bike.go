package bikes

// Coordinates of a specific location.
type Coordinates struct {
	Lat float64 `json:"latitude"  validate:"latitude,required"`
	Lng float64 `json:"longitude" validate:"longitude,required"`
}

// BikeLocationUpdate is sent when the position of a bike has changed.
type BikeLocationUpdate struct {
	Location Coordinates `json:"location" validate:"required"`
}

// BikeResponse is the response send.
type BikeResponse struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Location     Coordinates `json:"location"`
	Rented       bool        `json:"rented"`
	RentedByUser bool        `json:"returnable"`
}

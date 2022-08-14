package bikes

type Coordinates struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

type BikeLocationUpdate struct {
	Location Coordinates `json:"location"`
}

type BikeResponse struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Location     Coordinates `json:"location"`
	Rented       bool        `json:"rented"`
	RentedByUser bool        `json:"returnable"`
}

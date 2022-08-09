package bikes

type Coordinates struct {
	Lat string `json:"latitude"`
	Lng string `json:"longitude"`
}

type BikeResponse struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Location     Coordinates `json:"location"`
	Rented       bool        `json:"available"`
	RentedByUser bool        `json:"returnable"`
}

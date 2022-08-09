package rents

type Coordinates struct {
	Lat string `json:"latitude"`
	Lng string `json:"longitude"`
}

type RentRequest struct {
	BikeID        string      `json:"bikeID"`
	Renter        string      `json:"userID"`
	StartTime     string      `json:"startTime"`
	StartLocation Coordinates `json:"startLocation"`
}

type RentResponse struct {
	ID            string      `json:"id"`
	BikeID        string      `json:"bikeID"`
	Renter        string      `json:"userID"`
	StartTime     string      `json:"startTime"`
	EndTime       string      `json:"endTime"`
	StartLocation Coordinates `json:"startLocation"`
	EndLocation   Coordinates `json:"endLocation"`
}

type EndRentRequest struct {
	EndLocation Coordinates `json:"endLocation"`
}

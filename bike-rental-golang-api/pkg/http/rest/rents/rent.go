package rents

import "github.com/google/uuid"

type Coordinates struct {
	Lat string `json:"latitude"`
	Lng string `json:"longitude"`
}

type StartRequest struct {
	BikeID string `json:"bikeID"`
	Renter string `json:"userID"`
}

type StopRequest struct {
	Status uint8     `json:"status"`
	UserID uuid.UUID `json:"-"`
	RentID string    `json:"-"`
}

type RentResponse struct {
	ID            string      `json:"id"`
	BikeID        string      `json:"bikeID"`
	Renter        string      `json:"userID"`
	Status        uint8       `json:"status"`
	StartTime     string      `json:"startTime"`
	EndTime       string      `json:"endTime"`
	StartLocation Coordinates `json:"startLocation"`
	EndLocation   Coordinates `json:"endLocation"`
}

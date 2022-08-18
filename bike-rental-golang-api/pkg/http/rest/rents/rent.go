package rents

import "github.com/google/uuid"

// Coordinates of the location for start and stop.
// They are only used for view, therefor string is used as type.
type Coordinates struct {
	Lat string `json:"latitude"`
	Lng string `json:"longitude"`
}

// StartRequest is request to start a rent.
type StartRequest struct {
	BikeID string `json:"bikeID" validate:"uuid4,required"`
	Renter string `json:"userID" validate:"uuid4,required"`
}

// StopRequest is the request to stop a rent.
// The request body must only contain status with value 2(=Finished).
// UserID and RentID are provided by the handler.
type StopRequest struct {
	Status uint8     `json:"status" validate:"eq=2,required"`
	UserID uuid.UUID `json:"-" validate:"required"`
	RentID string    `json:"-" validate:"uuid4,required"`
}

// GetRentRequest is the request to retrieve the rents of a user.
// Values of the struct are provided by a query and need to be set.
type GetRentRequest struct {
	Status string `validate:"number,eq=1,required"`
	UserID string `validate:"uuid4,required"`
}

// RentResponse is the response send.
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

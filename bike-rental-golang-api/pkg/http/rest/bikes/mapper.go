package bikes

import (
	"github.com/nicoflink/bike-rental/pkg/list"
)

func mapToBikesJsonResponse(bikes []list.Bike) []BikeResponse {
	jBikes := make([]BikeResponse, 0, len(bikes))

	for _, b := range bikes {
		jBikes = append(jBikes, mapToBikeJsonResponse(b))
	}

	return jBikes
}

func mapToBikeJsonResponse(bike list.Bike) BikeResponse {
	return BikeResponse{
		ID:   bike.ID.String(),
		Name: bike.Name,
		Location: Coordinates{
			Lat: bike.Location.Lat,
			Lng: bike.Location.Lng,
		},
		Rented:       bike.Rented,
		RentedByUser: bike.RentedByUser,
	}
}

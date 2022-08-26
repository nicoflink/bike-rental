# bike-rental
Sample project for bike rental

# Run the app locally 
In order to run the app the rest-api server as well as the angular needs to be started. Moreover an API-Key for Google
Maps is required. The application is tested with:
- node v18.7.0
- npm v8.15.0
- go 1.16

## Start the REST-API Server

- `cd bike-rental-golang-api`
- `go run cmd/bike-rental-server/main.go`

The http server starts on `localhost:8080`

## Start the angular application 
In order to start the application insert your Google Maps API key into bike-map-app/.env

- `cd bike-map-app`
- `npm install`
- `npm start`

## Use the application 
Open a browser on `http://localhost:4200/` and you will see a map of Frankfurt/Main with gray and red markers. The gray
markers are bikes which are already rented. The red ones are available for rent. 

- Click on any red marker to see instructions
- If you want to rent the bike, click on the "Rent Bike" button
- The bike lock will be unlocked automatically and the marker becomes draggable - the marker turns gray
- Drag the marker to any position on the map to simulate a ride
- If you want to finish the rent, click on the marker again to open the info window and click on the "Return Bike" button - the marker turns red again
- In the console of the browser you will find a summary of you rent

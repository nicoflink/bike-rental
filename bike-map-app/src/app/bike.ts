// Reprensentation of a bike in the fronend
export interface Bike {
    id: string,
    name: string,
    location: {
        latitude: number,
        longitude: number 
    },
    rented: boolean,
    returnable: boolean // returnable is true if renter of the bike is the current user
}

// Information needed to update the location of the bike.
export interface BikeLocationUpdate {
    bikeID: string,
    location: {
        latitude: number,
        longitude: number 
    },
}
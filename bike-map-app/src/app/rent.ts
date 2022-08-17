// location to be shown in the bike details
interface Location {
    latitude: string,
    longitude: string
}

// Reprensentation of a rent in the fronend
export interface Rent {
    id: string,
    bikeID: string,
    userID: string,
    status: number,             // Status can be 0=Unkown, 1=Started or 2=Finished
    startTime: string,
    endTime: string,            // is empty in case status is Started (=1)
    startLocation: Location,
    endLocation: Location       // is empty in case status is Started (=1)
}
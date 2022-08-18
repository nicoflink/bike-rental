import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Bike, BikeLocationUpdate } from '../bike';

/**
 * BikeService provides functions to use the bikes backend API.
 */
@Injectable({
  providedIn: 'root'
})
export class BikeService {

  private route: string = "/bikes"

  constructor(private http: HttpClient) { }

  /**
   * getBikes returns all bikes stored in the backend. If no bikes are stored an empty bike is returned.
   * @returns List of all bikes
   */
  getBikes(): Observable<Bike[]> {
    return this.http.get<Bike[]>(this.route);
  }

  /**
   * updateLocation is used to update the location of the current bike after it has been moved (dragging the marker).
   * @param locationUpdate
   * @returns updated Bike
   */
  updateLocation(locationUpdate: BikeLocationUpdate):Observable<Bike> {
    return this.http.patch<Bike>(`${this.route}/${locationUpdate.bikeID}`, {
      "location": locationUpdate.location
    });
  }
}

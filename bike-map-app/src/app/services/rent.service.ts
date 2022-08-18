import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Rent } from '../rent';
import { SessionService } from './session.service';

/**
 * RentService provides functions to use the rents backend API.
 */
@Injectable({
  providedIn: 'root'
})
export class RentService {

  private route: string = "/rents"

  constructor(
    private http: HttpClient,
    private session: SessionService) { }
  
  /**
   * startRent creates a rental in the backend by providing the bike and the user of the session. 
   * @param bikeID 
   * @returns Observable of Rent that has been created 
   */
  startRent(bikeID: string):Observable<Rent> {
    const userID = this.session.getSessionID();
    
    return this.http.post<Rent>(this.route, { bikeID, userID });
  }

  /**
   * stopRent finished the rent of the bike in the backend by updating the status of the rent to 2 (=Finished). 
   * @param rentID 
   * @returns Observable of Rent that provides a summary of the rent. 
   */
  stopRent(rentID: string):Observable<Rent> {
    return this.http.patch<Rent>(`${this.route}/${rentID}`, { status: 2 });
  }

  /**
   * getStartedRents returns open (status=1 Started) rents of a user.  
   * @returns 
   */
  getStartedRents():Observable<Rent[]> {
    const sessionID = this.session.getSessionID();

    return this.http.get<Rent[]>(`${this.route}?status=1&userID=${sessionID}`);
  }
}

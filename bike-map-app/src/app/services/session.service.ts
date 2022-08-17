import { Injectable } from '@angular/core';
import { v4 as uuidv4 } from 'uuid';

/**
 * SessionService is used to create a session to differentiate users in the backend. For this a generated uuid is stored in the session storage of the browser.
 */
@Injectable({
  providedIn: 'root'
})
export class SessionService {

  // Key for the session that is used in the session storage. 
  sessionKey: string = "session" 

  constructor() { }

  // generateSession generates a new uuid und stores it in the session storage.
  generateSession(): string {
    const sessionID = uuidv4();
    sessionStorage.setItem(this.sessionKey, sessionID)

    return sessionID;
  }

  // getSessionID retrieves the session ID from the session storage. If no seesion ID is available, a new one is generated.
  getSessionID(): string {
    let sessionID = sessionStorage.getItem(this.sessionKey);

    if (!sessionID){
      sessionID = this.generateSession();  
    }

    return sessionID;
  }
}

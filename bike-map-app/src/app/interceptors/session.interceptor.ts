import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { SessionService } from '../services/session.service';

/**
 * SessionInterceptor sets the http header "SessionID" for the current session to differentiate single session (here used for users) in the backend.
 */
@Injectable()
export class SessionInterceptor implements HttpInterceptor {

  constructor(private sessionService: SessionService) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    const SessionID = this.sessionService.getSessionID();

    return next.handle(request.clone({setHeaders: {SessionID}}));
  }
}

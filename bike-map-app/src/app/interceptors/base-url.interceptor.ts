import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';

/**
 * BaseURLInterceptor is used to the set the base URL path (api prefix) defined in the environment. 
 */
@Injectable()
export class BaseURLInterceptor implements HttpInterceptor {

  // basePath to be used
  private basePath: string = "/api/v1";

  constructor() {} 

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    const apiReq = request.clone({ url: `${this.basePath}${request.url}` });
    
    return next.handle(apiReq);
  }
}

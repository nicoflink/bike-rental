import { NgModule } from '@angular/core';
import { GoogleMapsModule } from '@angular/google-maps';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS, HttpClientJsonpModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { BaseURLInterceptor } from './interceptors/base-url.interceptor';
import { SessionInterceptor } from './interceptors/session.interceptor';
import { Bike2MarkerPipe } from './pipes/bike2-marker.pipe';
import { BikeDetailsComponent } from './bike-details/bike-details.component';
import { MatButtonModule } from '@angular/material/button';
import { FlexLayoutModule } from '@angular/flex-layout';

@NgModule({
  declarations: [
    AppComponent,
    Bike2MarkerPipe,
    BikeDetailsComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    HttpClientJsonpModule,
    GoogleMapsModule,
    MatButtonModule,
    FlexLayoutModule
  ],
  providers: [
    {provide: HTTP_INTERCEPTORS, useClass: BaseURLInterceptor, multi: true},
    {provide: HTTP_INTERCEPTORS, useClass: SessionInterceptor, multi: true},
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

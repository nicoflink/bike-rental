import { Component } from '@angular/core';
import { Bike } from './bike';
import { BikeService } from './services/bike.service';
import { MapInfoWindow, MapMarker } from '@angular/google-maps';
import { ViewChild } from '@angular/core';
import { catchError, map, Observable, of } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';

/**
 * AppComponent is the main component of the application which is the map showing the bikes of the rental service.
 */
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less']
})
export class AppComponent {
  title = 'bike-map-app';

  // used for lazy load of the Google Maps JavaScript API.
  apiLoaded: Observable<boolean>;

  // Reference to the bike details info window - it is used to open and close the window.
  @ViewChild(MapInfoWindow) infoWindow!: MapInfoWindow;

  // Are the bikes to be shown on the map. Bikes are always set.
  bikes!: Bike[]

  // Is the bike selected via marker
  selectedBike?: Bike;

  // currentRent is set to true, if a bike is already rented by the user
  currentRent: boolean = false;

  constructor(
    private bikeService: BikeService,
    private httpClient: HttpClient){
      this.apiLoaded = this.httpClient.jsonp(`https://maps.googleapis.com/maps/api/js?key=${environment.apiKey}`, 'callback')
        .pipe(
          map(() => true),
          catchError(() => of(false)),
        );
    }

  // hardcoded starting point for POC - Frankfurt
  center: google.maps.LatLngLiteral = {lat: 50.10737359920496, lng: 8.667990906890859};
  zoom = 15

  ngOnInit() {
    /*
    navigator.geolocation.getCurrentPosition((position) => {
      this.center = {
        lat: position.coords.latitude,
        lng: position.coords.longitude,
      }
    })
    */

    this.getBikes();
  }

  /**
   *  getBikes retrieves all bike to be shown on the map. If a bike is returnable, the user has a started rent with this bike.
   */
  getBikes() {
    this.bikeService.getBikes()
    .subscribe(
      bikes => {
        this.bikes = bikes;
        // iterate over bikes to find possible rent of current user
        bikes.forEach(bike => {
          if (bike.returnable) {
            this.currentRent = true;
          }
        });

        if(this.infoWindow) {
          this.infoWindow.close();
        }
      });
  }

  /**
   * handleBikeChange handles changes comming from bike details info window when the user started or stopped a rent.
   * @param bike updated bike
   */
  handleBikeChange(bike: Bike): void {
    this.currentRent = bike.returnable;
    
    this.bikes.forEach(function(b, index, bikes){
      if(b.id == bike.id){
        bikes[index] = bike;
      }
    });

    if(this.infoWindow) {
      this.infoWindow.close();
    }
  }

  /**
   * trackByRentStatus is used to rerender the map in case of changes through a rent. As trigger for changes the attribute rented of 
   * the bike is used.
   * @param index 
   * @param b 
   * @returns 
   */
  trackByRentStatus(index: number, b: Bike): boolean {
    return b.rented;
  }

  /**
   * updateBikeLocation updated the location of a bike in the backend when the marker has been dragged 
   * @param event
   * @param bikeID 
   * @returns 
   */
  updateBikeLocation(event:google.maps.MapMouseEvent, bikeID: string): void {
    if(!event.latLng){
      console.error("latLng undefined during location update");
      
      return;
    }

    this.bikeService.updateLocation({
      bikeID: bikeID,
      location: {
        latitude: event.latLng.lat(),
        longitude: event.latLng.lng()
      }
    }).subscribe(updatedBike => {
      if (this.selectedBike){
        this.selectedBike!.location = updatedBike.location;
      }});
  }

  /**
   * openInfoWindow opens the info window with bike details and sets selectedBike.
   * @param marker 
   * @param bike 
   */
  openInfoWindow(marker:MapMarker, bike: Bike) {
    this.selectedBike = bike;
    this.infoWindow.open(marker);
  }
}

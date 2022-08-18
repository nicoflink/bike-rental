import { Pipe, PipeTransform } from '@angular/core';
import { Bike } from '../bike';

/**
 * Bike2MarkerPipe is used to convert a bike object into a marker as specified in google.maps.MarkerOptions.
 */
@Pipe({
  name: 'bike2Marker'
})
export class Bike2MarkerPipe implements PipeTransform {

  private gray: string = "A7A19F"
  private red: string = "DD2902"

  /**
   * transform converts a bike into a google maps marker. 
   * For this, the position of the bike is used for the marker.
   * In case the bike is rented by the current user (returnable) the markers becomes draggable to simulate a ride with the bike.
   * The icon color specifies whether a bike is available (red) or already rented (gray).
   * @param bike 
   * @param args 
   * @returns Marker for the map as specified in google.maps.MarkerOptions
   */
  transform(bike: Bike, ...args: unknown[]): google.maps.MarkerOptions {
    return {
      position: {
        lat: bike.location.latitude,
        lng: bike.location.longitude
      },
      draggable: bike.returnable,
      animation: bike.returnable? google.maps.Animation.DROP: null,
      icon: bike.rented? this.getGoogleIcon(this.gray): this.getGoogleIcon(this.red),
    }
  }

  private getGoogleIcon(color:string): string {
    return `https://mt.google.com/vt/icon/name=icons/onion/SHARED-mymaps-pin-container-bg_4x.png,icons/onion/SHARED-mymaps-pin-container_4x.png,icons/onion/1899-blank-shape_pin_4x.png&highlight=ff000000,${color},ff000000&scale=1.7`
  }
}

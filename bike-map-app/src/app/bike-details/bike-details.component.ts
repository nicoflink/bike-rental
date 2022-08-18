
import { Component, Input, Output, EventEmitter, OnChanges, SimpleChanges } from '@angular/core';
import { tap } from 'rxjs';
import { Bike } from '../bike';
import { Rent } from '../rent';
import { RentService } from '../services/rent.service';

/**
 * BikeDetailsComponent is the popup opening after clicking the a marker in the map. It shows some details about the current rent, the status of the rent
 * and enables the user to start or stop the rent. 
 */
@Component({
  selector: 'app-bike-details',
  templateUrl: './bike-details.component.html',
  styleUrls: ['./bike-details.component.less']
})
export class BikeDetailsComponent implements OnChanges {

  // Bike for which the details are supposed to be shown
  @Input() bike: Bike | undefined;
  // Specifies whether a user has already rented a bike 
  @Input() hasAlreadyRent: boolean = true

  // rentChange emits an event after the start or stop of the rent - returns the updated bike to update the map
  @Output() rentChange = new EventEmitter<Bike>();
  // forceReload emits an event to force the rendering of the map again, e.g. if an error occurs
  @Output() forceReload = new EventEmitter();

  // are the details for a rent - can be undefined in case no rent is started
  rentDetails?: Rent

  constructor(private rentService: RentService) { }

  ngOnChanges(changes: SimpleChanges): void {
    if (this.hasAlreadyRent && !this.rentDetails) {
      this.getCurrentRent()
    }
  }

  /**
   * getCurrentRent gets the current rent of a user and sets the variable rentDetails. 
   * A user is allowed to rent a single bike at the same time.
   * In case the service returns multiple rents a warning is logged.
   */
  getCurrentRent(): void {
    this.rentService.getStartedRents().subscribe(rents => {
      if (rents.length > 0) {
        this.rentDetails = rents[0];
      }

      if (rents.length > 1) {
        console.warn("Received more than 1 rent");
      }
    })
  }

  /**
   *  startRent starts renting of the selected bike for the current user. rentService provides details about
   *  the rent which are set to rentDetails. By starting the rent, the bike status changes to rented and returnable.
   *  In case of an error, the error is logged and the map will be rendered again.
   */
  startRent(): void {
    if (!this.bike) {
      console.error("bike is not defined");
      return;
    }

    this.rentService.startRent(this.bike.id)
      .pipe(
        tap(() => this.hasAlreadyRent = true),
        tap(() => {
          this.bike!.rented = true;
          this.bike!.returnable = true;
        })
      )
      .subscribe({
        next: rent => {
          this.rentDetails = rent;
          this.rentChange.emit(this.bike);
        },
        error: e => {
          console.error("Start Rent:", e);
          this.forceReload.emit();
        }
      });
  }

  /**
   *  stopRent stops renting of the selected bike for the current user. rentService provides a summary of the rent, that is logged to the console. 
   *  After finishing the rent, the bike status changes to not rented and not returnable.
   */
  stopRent(): void {
    if (!this.bike) {
      console.error("bike is not defined");
      return;
    }

    if (!this.rentDetails) {
      console.error("rent details are not defined");
      return;
    }

    this.rentService.stopRent(this.rentDetails.id)
      .pipe(
        tap(() => this.hasAlreadyRent = false),
        tap(() => {
          this.bike!.rented = false;
          this.bike!.returnable = false;
        })
      )
      .subscribe(rent => {
        console.log("Thanks for renting our bike!", "Rent summary:", rent);
        this.rentDetails = undefined;
        this.rentChange.emit(this.bike);
      });
  }
}

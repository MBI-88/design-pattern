package main

import (
	"fmt"
	"factory/cmd"
)



func main() {
	reservation := cmd.NewReservationFactory("hotel", "20231202")
	trip := cmd.NewTrip()
	fmt.Printf("Reservation done: %s ",reservation.GetReservationDate() )
	trip.AddReservation(reservation)
}
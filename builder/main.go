package main

import (
	"builder/cmd"
	"fmt"
)





func main() {
	builder := cmd.NewReservationBuilder()
	builder.ReservationDate("20231210")
	flight := builder.Vertical("flight").Build()
	
	fmt.Printf("Cancelation date %f\n", flight.CalculateCancellationFee())

	fmt.Printf("Reservation date %s\n", flight.GetReservationDate())

	hotel := builder.Vertical("hotel").Build()

	fmt.Printf("Cancelation date %f\n", hotel.CalculateCancellationFee())

	fmt.Printf("Reservation date %s\n", hotel.GetReservationDate())
	
	
}
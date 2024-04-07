package main

import (
	"abstract/cmd"
	"fmt"
)




func main() {
	hotelFactory := cmd.NewAbstractFactory("hotel")
	reservation := hotelFactory.CreateReservation()
	invoice := hotelFactory.CreateInvoice()

	reservation.SetReservationDate("20231201")

	fmt.Printf("Reservation date %s\n", reservation.GetReservationDate())
	invoice.SetInvoce(200.50, "20231216")
	fmt.Printf("Invoice %f\n", invoice.CalculateInvoce())

	flightFactory := cmd.NewAbstractFactory("flight")
	reservation = flightFactory.CreateReservation()
	invoice = flightFactory.CreateInvoice()

	reservation.SetReservationDate("20230110")
	fmt.Printf("Reservation date %s\n", reservation.GetReservationDate())
	invoice.SetInvoce(800.20, "20230110")
	fmt.Printf("Invoice %f\n", invoice.CalculateInvoce())
}
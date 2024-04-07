package main

import "bridge/cmd"

func main() {
	res := cmd.NewReservation("normal")
	res.CreateReservation("John", "Doe", 784598123)
	res.Cancel()

	premium := cmd.NewReservation("premium")
	premium.CreateReservation("Maria", "Smith", 345123698)
	premium.Cancel()

}

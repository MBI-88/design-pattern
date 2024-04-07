package cmd

import "fmt"

type InstitutionSeller struct {
	name     string
	lastname string
	dni      int
}

func (i InstitutionSeller) CancelReservation(charge float64) {
	fmt.Printf("Institution seller %s %s canceled reservation change = %f\n", i.name, i.lastname, charge)
}

func (i *InstitutionSeller) SetData(name, lastname string, dni int) {
	i.dni = dni
	i.lastname = lastname
	i.name = name
}

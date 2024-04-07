package cmd

import "fmt"

type SmallScaleSeller struct {
	name     string
	lastname string
	dni      int
}

func (s SmallScaleSeller) CancelReservation(charge float64) {
	fmt.Printf("Normal seller %s %s canceled reservation change = %f\n", s.name, s.lastname, charge)
}

func (s *SmallScaleSeller) SetData(name, lastname string, dni int) {
	s.dni = dni
	s.lastname = lastname
	s.name = name
}

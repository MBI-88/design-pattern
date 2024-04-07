package cmd

import "fmt"

type contextA struct {
	status *context
}

func (ca *contextA) Op1(c *context) {
	ca.status = c
	fmt.Printf("Operation 1 data: %s %t  %t \n", ca.status.name, ca.status.chacked,ca.status.successful)
}

func (ca *contextA) Op2(c *context) {
	ca.status = c
	fmt.Printf("Operation 2 data: %s %t  %t \n", ca.status.name, ca.status.chacked,ca.status.successful)
}


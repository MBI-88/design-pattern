package cmd

import "fmt"


// adapteeToArray example
type adapteeToArray struct {
	name     string
	lastname string
	id       int
}

// Sets parameter needed
func (a *adapteeToArray) SetParameters(name, lastname string, id int) {
	a.name = name
	a.lastname = lastname
	a.id = id
}

// Returns an array
func (a adapteeToArray) GeParams() []string {
	var arr []string
	arr = append(arr, a.name, a.lastname, fmt.Sprintf("%d", a.id))
	return arr
}

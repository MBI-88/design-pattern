package cmd

import "fmt"

// adapteeToMap example
type adapteeToMap struct {
	name     string
	lastname string
	id       int
}

// Sets parameter needed
func (a *adapteeToMap) SetParameters(name, lastname string, id int) {
	a.name = name
	a.lastname = lastname
	a.id = id
}

// Returns a map
func (a adapteeToMap) GeParams() map[string]string {
	mapAdptee := make(map[string]string)
	mapAdptee["name"] = a.name
	mapAdptee["lastname"] = a.lastname
	mapAdptee["id"] = fmt.Sprintf("%d", a.id)
	return mapAdptee
}

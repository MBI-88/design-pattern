package cmd

import (
	"fmt"
	"strings"
)

type creditCard struct {
	name string
	lastname string
	dni string
	id int
	status string
	credit int

}

func (c *creditCard) CreateCreditCard(name,lastname,dni string, id int) {
	c.name = strings.Trim(name, " ")
	c.lastname = strings.Trim(lastname, " ")
	c.dni = strings.Trim(dni, " ")
	c.id = id
	c.credit = 50000
	c.status = "activated"
}

func (c *creditCard) CancelCreditCard() {
	c.status = "cancelled"
}

func (c creditCard) GetInfo() map[string]string {
	data := make(map[string]string)
	data["name"] = c.name
	data["lastname"] = c.lastname
	data["dni"] = c.dni
	data["status"] = c.status
	data["id"] = fmt.Sprintf("%d", c.id)

	return data
}
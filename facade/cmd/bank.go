package cmd

import (
	"fmt"
	"strings"
)


type bankAccount struct {
	name string
	lastname string
	dni string
	amount float64
	status string
	id int
}


func (b *bankAccount) CreateAccount(name,lastname,dni string, amount float64 ) {
	b.name = strings.Trim(name, " ")
	b.lastname = strings.Trim(lastname, " ")
	b.dni = strings.Trim(dni, " ")
	b.amount = amount
	b.status = "reserved"
}

func (b *bankAccount) CancelAccount() {
	b.status = "canelled"
}

func (b bankAccount) GetInfo() map[string]string {
	data := make(map[string]string)
	data["name"] = b.name
	data["lastname"] = b.lastname
	data["dni"] = b.dni
	data["amount"] = fmt.Sprintf("%f", b.amount)
	data["status"] = b.status

	return data
}
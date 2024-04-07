package cmd

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	u    *user
)

type user struct {
	name     string
	lastname string
	address  string
	dni      int
}

func (u *user) SetUser(name, lastname, address string, dni int) {
	u.address = address
	u.dni = dni
	u.lastname = lastname
	u.name = name
}

func (u user) GetUser() map[string]string {
	us := make(map[string]string)
	us["name"] = u.name
	us["lastname"] = u.lastname
	us["address"] = u.address
	us["dni"] = fmt.Sprintf("%d", u.dni)

	return us
}

package cmd

import (
	"fmt"
	"math/rand"
	"strings"
)

type facade struct {
	card    creditCardInt
	account banckInt
}

func (f *facade) CreateUser(name, lastname, dni string, amount float64) string {
	f.card = new(creditCard)
	f.account = new(bankAccount)
	name = strings.Trim(name, " ")
	lastname = strings.Trim(lastname, " ")
	dni = strings.Trim(dni, " ")
	id := rand.Intn(100000000) + 100000000
	if amount < 5000 {
		return fmt.Sprintf("Your amount is under 5000$")
	}
	f.card.CreateCreditCard(name, lastname, dni, id)
	f.account.CreateAccount(name, lastname, dni, amount)
	return fmt.Sprintf("Count created successful!")
}

func (f facade) CancelUser() {
	f.account.CancelAccount()
	f.card.CancelCreditCard()
	fmt.Println("User's account was cancelled")
}

func (f facade) GetUserInfo() map[string]map[string]string {
	user := make(map[string]map[string]string)
	user["creditcard"] = f.card.GetInfo()
	user["bankAccount"] = f.account.GetInfo()
	return user
}

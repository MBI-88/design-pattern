package main

import (
	"facade/cmd"
	"fmt"
)

func showInfo(data map[string]map[string]string) {
	fmt.Println("Account status")
	fmt.Println("{")
	for key, val := range data {
		fmt.Printf(" %s: {\n", key)
		for k, v := range val {
			fmt.Printf("\t%s: %s\n",k, v)
		}
		fmt.Println("}")
	}
	fmt.Println("}")
}

func main() {
	user := cmd.NewAccount()
	fmt.Println(user.CreateUser("John", "Doe","12345678A", 20000))

	showInfo(user.GetUserInfo())

	user.CancelUser()

	showInfo(user.GetUserInfo())

	
}
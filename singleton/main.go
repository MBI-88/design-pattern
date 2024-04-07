package main

import (
	"fmt"
	"singleton/cmd"
)

func main() {
	us1 := cmd.NewUser()
	us1.SetUser("John", "Do", "Avenue 8 Miami FL", 45857595)

	fmt.Printf("User %v\n", us1.GetUser())

	us2 := cmd.NewUser()

	fmt.Printf("User %v\n", us2.GetUser())

}

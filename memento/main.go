package main

import (
	"fmt"
	"memento/cmd"
)

func main(){
	ori := cmd.Creataker()
	ori.SetState("Activated")

	fmt.Printf("Actual state = %s\n", ori.GetState())

	mem := ori.GetMemento()
	ori.SetState("Deactivated")
	fmt.Printf("Actual state = %s\n", ori.GetState())

	ori.Restore(mem)
	fmt.Printf("Restored state = %s\n", ori.GetState())

}
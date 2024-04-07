package main

import "command/cmd"

func main() {
	commadB := cmd.NewCommand("commandB")
	commandA := cmd.NewCommand("commandA")

	commandA.Run()
	commadB.Run()
}
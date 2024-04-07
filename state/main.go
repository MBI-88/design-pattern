package main

import (
	"state/cmd"
)


func main() {
	context := cmd.NewContex()

	context.Op1()
	context.Op2()
	context.Op1()
	context.Op2()

}
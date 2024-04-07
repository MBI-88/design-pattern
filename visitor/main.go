package main

import (
	"visitor/cmd"
)


func main() {
	aggregate := []cmd.Node{cmd.NewNode("nodeA"),cmd.NewNode("nodeB")}
	visitor := cmd.NewVistior()

	for _, node := range aggregate {
		node.Accept(visitor)
	}
}
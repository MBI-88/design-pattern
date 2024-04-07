package cmd

import (
	"fmt"
)

type visitor struct {}

func (v visitor) Visit(node Node) {
	fmt.Println("Checking messages")
	switch node.(type) {
	case nodeA:
		fmt.Printf("Node A message: %s\n", node.GetMessage())
	case nodeB:
		fmt.Printf("Node B message: %s\n", node.GetMessage())
	}
}

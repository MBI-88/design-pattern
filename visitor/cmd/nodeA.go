package cmd

import "fmt"

type nodeA struct {
	message string
}

func (a nodeA) Accept(visitor Visitor) {
	fmt.Println("Visitor being visited node A")
	a.setMessage("visited")
	visitor.Visit(a)
}

func (a *nodeA) setMessage(m string) {
	a.message = m
}

func (a nodeA) GetMessage() string {
	return a.message
}
package cmd

import "fmt"

type nodeB struct {
	message string
}

func (b nodeB) Accept(visitor Visitor) {
	fmt.Println("Visitor being visited node B")
	b.setMessage("visited")
	visitor.Visit(b)
}

func (b *nodeB) setMessage(m string) {
	b.message = m
}

func (b nodeB) GetMessage() string {
	return b.message
}
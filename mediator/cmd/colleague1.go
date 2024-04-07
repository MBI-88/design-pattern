package cmd

import "fmt"

type Colleague1 struct {
	mediator MediatorInt
	state    string
}

func (c *Colleague1) SetMediator(mediator MediatorInt) {
	c.mediator = mediator
}

func (c *Colleague1) SetState(state string) {
	c.state = state
	fmt.Println("Colleague1: setting state: ", state)
}

func (c Colleague1) GetState() string {
	return c.state
}
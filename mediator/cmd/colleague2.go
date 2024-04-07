package cmd

import "fmt"

type Colleague2 struct {
	mediator MediatorInt
	state    int16
}

func (c *Colleague2) SetState(state int16) {
	c.state = state
	fmt.Println("Colleague2: setting state: ", state)
}

func (c Colleague2) GetState() int16 {
	return c.state
}

func (c *Colleague2) SetMediator(mediator MediatorInt) {
	c.mediator = mediator
}
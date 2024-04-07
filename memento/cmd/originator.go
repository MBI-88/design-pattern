package cmd

import "fmt"

type originator struct {
	state string
}

func (o originator) GetState() string {
	return o.state
}

func (o *originator) SetState(state string) {
	o.state = state
	fmt.Printf("Setting state to %s\n", o.state )
}

func (o originator) GetMemento() memento {
	return memento{o.state}
}

func (o *originator) Restore(me memento) {
	o.state  = me.GetState()
}
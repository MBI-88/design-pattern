package cmd

import "fmt"

type observerA struct {
	sub   SubInt
	state string
}

func (a *observerA) Update() {
	a.state = a.sub.GetState()
	fmt.Printf("Observer A: updated view state to: %s\n", a.state)
}

func (a *observerA) SetModel(s SubInt) {
	a.sub = s
	fmt.Println(a.sub.GetState())
}

type observerB struct {
	sub   SubInt
	state string
}

func (b *observerB) Update() {
	b.state = b.sub.GetState()
	fmt.Printf("Observer B: updated view state to: %s\n", b.state)
}

func (b *observerB) SetModel(s SubInt) {
	b.sub = s
}

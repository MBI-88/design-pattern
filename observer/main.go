package main

import (
	"fmt"
	"observer/cmd"
)


func main() {
	subject := cmd.NewSubject()
	observerA := cmd.NewObserver("A")
	observerB := cmd.NewObserver("B")

	observerA.SetModel(subject)
	observerB.SetModel(subject)

	subject.Attach(observerA)
	subject.Attach(observerB) 

	for st := range 4 {
		subject.SetState(fmt.Sprintf("state-%d", st))
	}	


}
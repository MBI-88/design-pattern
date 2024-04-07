package main


import (
	"mediator/cmd"
)


func main() {
	mediator := cmd.NewMediator()
	c1 := new(cmd.Colleague1)
	c2 := new(cmd.Colleague2)

	mediator.SetColleagueC1(c1)
	mediator.SetColleagueC2(c2)
	mediator.SetState("20")
}
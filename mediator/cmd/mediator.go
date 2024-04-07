package cmd

import "strconv"

type mediator struct {
	c1 *Colleague1
	c2 *Colleague2
}

func (m *mediator) SetColleagueC1(c1 *Colleague1) {
	m.c1 = c1
}

func (m *mediator) SetColleagueC2(c2 *Colleague2) {
	m.c2 = c2
}

func (m *mediator) SetState(s string) {
	m.c1.SetState(s)
	stateAsString, er := strconv.Atoi(s)
	if er == nil {
		m.c2.SetState(int16(stateAsString))
	}
}

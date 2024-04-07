package cmd

type subject struct {
	sub   []ObserverInt
	state string
}

func (s *subject) Attach(ob ObserverInt) {
	s.sub = append(s.sub, ob)
}

func (s subject) GetState() string {
	return s.state
}

func (s *subject) SetState(state string) {
	s.state = state
	for _, ob := range s.sub {
		ob.Update()
	}
}

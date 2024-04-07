package cmd

type SubInt interface {
	SetState(state string)
	GetState() string
	Attach(ob ObserverInt)
}

type ObserverInt interface {
	Update()
	SetModel(s SubInt)
}

func NewSubject() SubInt {
	return &subject{}
}

func NewObserver(ob string) ObserverInt {
	switch ob {
	case "B":
		return &observerB{}
	default:
		return &observerA{}
	}

}

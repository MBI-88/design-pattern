package cmd




type Func func(age int)

type Casino interface {
	Setcasino(name, gender string, age int)
	Getcasino() map[string]string
}

func NewCasino() Casino {
	return &casino{}
}

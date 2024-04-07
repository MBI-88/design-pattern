package cmd

import "fmt"

type casino struct {
	name   string
	age    int
	gender string
	status bool
}

func (u *casino) Setcasino(name, gender string, age int) {
	u.age = age
	u.name = name
	u.gender = gender
	u.status = false
}

func (u casino) Getcasino() map[string]string {
	userAge := u.Decorator(u.ValidAge)
	userAge(u.age)
	dict := make(map[string]string)
	dict["name"] = u.name
	dict["age"] = fmt.Sprintf("%d", u.age)
	dict["gender"] = u.gender
	dict["status"] = fmt.Sprintf("%t", u.status)
	return dict
}

func (u *casino) Decorator(fn Func) Func {
	return func(age int) {
		fn(age)
	}
}

func (u *casino) ValidAge(age int) {
	if age > 18 {
		u.status = true
	}
}

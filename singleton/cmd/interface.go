package cmd

type User interface {
	SetUser(name, lastname, address string, dni int)
	GetUser() map[string]string
}

func NewUser() User {

	once.Do(func() {
		u = &user{}
	})

	return u
}

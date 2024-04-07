package cmd

type creditCardInt interface {
	CreateCreditCard(name,lastname,dni string, id int)
	CancelCreditCard()
	GetInfo() map[string]string
}

type banckInt interface {
	CreateAccount(name,lastname,dni string, amount float64 )
	CancelAccount()
	GetInfo() map[string]string
}

type facadeInt interface {
	CreateUser(name,lastname,dni string, amount float64) string
	GetUserInfo() map[string]map[string]string
	CancelUser()
}



func NewAccount() facadeInt {
	return &facade{}
}

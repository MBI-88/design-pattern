package cmd



type commandA struct{
	re *receiver
}

func(c *commandA) Execute() {
	c.re.Action("Command A executed!")
}
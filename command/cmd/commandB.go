package cmd


type commandB struct {
	re *receiver
}

func(c *commandB) Execute() {
	c.re.Action("Command B exectued!")
}
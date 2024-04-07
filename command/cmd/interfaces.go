package cmd

type Report interface {
	Execute()
}

type InvokerInt interface {
	Run()
	Schedule(cmd Report)
}

func NewCommand(cmd string) InvokerInt {
	switch cmd {
	case "commmandB":
		commad := &commandB{&receiver{}}
		inst := new(invoker)
		inst.Schedule(commad)
		return inst
	default:
		commad := &commandA{&receiver{}}
		inst := new(invoker)
		inst.Schedule(commad)
		return inst
	}

}

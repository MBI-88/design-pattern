package cmd


type ChanInt interface {
	Finish() error
	SetNext( next *chainReceiver)
	Handle(what string) error
}

// Add commands to the Chain's object
func helper(inst *chainReceiver, command string) *chainReceiver {
	if inst.canHanle != "" {
		temp := new(chainReceiver)
		temp.canHanle = command
		temp.next = inst
		return temp
	}

	inst.canHanle = command
	return inst
}

// Create a new Chain Responsbility object
func NewChain() ChanInt {
	inst := new(chainReceiver)
	for _, command := range []string{"cd","pwd","ls","chmod"} {
		inst = helper(inst, command)
	}
	return inst
}


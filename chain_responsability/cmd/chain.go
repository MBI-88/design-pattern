package cmd

import (
	"errors"
	"fmt"
)

type chainReceiver struct {
	canHanle string
	next     *chainReceiver
}

func (r *chainReceiver) Finish() error {
	fmt.Println(r.canHanle, " Receiver Finishing")
	return nil
}

func (r *chainReceiver) SetNext(next *chainReceiver) {
	r.next = next
}

func (r *chainReceiver) Handle(what string) error {
	if what == r.canHanle {
		return r.Finish()
	} else if r.next != nil {
		return r.next.Handle(what)
	} else {
		fmt.Println("No Receiver could handle the request!")
		return errors.New("No Receiver to Handle")
	}
}

package main

import (
	"fmt"
	"strategy/cmd"
	"time"
)

func main(){
	var (
		start int
		stop int
	)
	sta := cmd.NewStrategy("a")
	stb := cmd.NewStrategy("b")

	start = time.Now().Nanosecond()
	fmt.Printf("Output %d\n",sta.FindBreadth([]int{1,12,3,4,7,9,2,8,30,25,45,78,91,65,11,2,6,10}))
	stop  = time.Now().Nanosecond()

	fmt.Printf("Time taken %d ns\n", stop - start)

	start = time.Now().Nanosecond()
	fmt.Printf("Output %d\n",stb.FindBreadth([]int{1,12,3,4,7,9,2,8,30,25,45,78,91,65,11,2,6,10}))
	stop  = time.Now().Nanosecond()

	fmt.Printf("Time taken %d ns\n", stop - start)
}
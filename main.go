package main

import (
	"fmt"
	"time"
)

type GoLimiter struct {
	fn              func()
	push            func(arg interface{})
	maximumRoutines int
}

var queue = make([]int, 0)

func goLimiter(limit int, isDone chan bool) {
	ch := make(chan bool, limit)
	counter := 0
	i := 0

	for i < len(queue) {
		for counter < limit && i < len(queue) {
			go ApiCall(queue[i], ch)
			fmt.Println("Go routine ", i, " ", len(queue))
			i++
			counter++
		}

		b := <-ch
		if b {
			counter--
		}
	}

	//isDone <- true
}

func ApiCall(i int, ch chan bool) {
	time.Sleep(time.Second * time.Duration(i))
	fmt.Println("Done with param", i)
	ch <- true
}

func pushToQueue(arg interface{}) {
	queue = queue.append(arg)
}

func NewGoRoutineLimiter(maximumRoutines int, processor func()) GoLimiter {
	var queue = make([]interface{}, 0)

	return GoLimiter{
		push:            pushToQueue,
		maximumRoutines: maximumRoutines,
		fn:              processor,
	}

}

func main() {
	isDone := make(chan bool)
	queue = append(queue, 3)
	queue = append(queue, 3)
	queue = append(queue, 2)

	goLimiter(1, isDone)
}

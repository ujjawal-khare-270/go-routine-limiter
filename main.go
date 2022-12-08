package main

import (
	"time"
)

var queue = make([]int, 1000)

func goLimiter(limit int, isDone chan bool) {
	ch := make(chan bool, limit)
	counter := 0

	for len(queue) > 0 {
		for counter < limit {
			go ApiCall(queue[0], ch)
			counter++

			b := <-ch
			if b {
				counter--
			}
			queue = append(queue[1:])
		}
	}

	isDone <- true
}

func ApiCall(i int, ch chan bool) {
	time.Sleep(time.Duration(i))
	ch <- true
}

func main() {
	isDone := make(chan bool)
	queue = append(queue, 1)
	goLimiter(2, isDone)
}

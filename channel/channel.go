package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c{
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	ch := make(chan int)
	go worker(id, ch)
	return ch
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

func closeChannel() {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	bufferedChannel()
	closeChannel()
}

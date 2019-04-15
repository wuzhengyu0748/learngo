package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for{
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func worker(id int, c chan int) {
	for n := range c{
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	ch := make(chan int)
	go worker(id, ch)
	return ch
}

func main() {
	var c1, c2 = generator(), generator() // c1 and c2 = nil
	var worker = createWorker(0)
	var values []int
	var activeValue int
	tm := time.After(10 * time.Second)
	tick := time.Tick(1 * time.Second)
	for{
		var activeWorker chan <- int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <- c1:
			values = append(values, n)
		case n := <- c2:
			values = append(values, n)
		case <- time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <- tick:
			fmt.Println("Queue len = ", len(values))
		case activeWorker <- activeValue:
			values = values[1:]
		case <- tm:
			fmt.Printf("bye~")
			return
		}
	}
}

package main

import "fmt"

func generator() chan int {
	out := make(chan int)
	go func() {

	}()
	return out
}

func main() {
	var c1, c2 chan int // c1 and c2 = nil
	select {
	case n := <- c1:
		fmt.Println("Received from c1:", n)
	case n := <- c2:
		fmt.Println("Received from c2:", n)
	default:
		fmt.Println("No value received")
	}
}

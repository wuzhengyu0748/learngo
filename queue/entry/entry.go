package main

import (
	"wuzhengyu/learngo/queue"
	"fmt"
)

func main() {
	myQueue := queue.Queue{1}

	myQueue.Push(2)
	myQueue.Push(3)

	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.IsEmpty())
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.IsEmpty())

	myQueue.Push("abc")
	fmt.Println(myQueue.Pop())
}

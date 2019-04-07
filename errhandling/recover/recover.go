package main

import (
	"fmt"
	"errors"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(err)
		}
	}()
	//panic(123)
	panic(errors.New("this is an error"))
}

func main() {
	tryRecover()
}

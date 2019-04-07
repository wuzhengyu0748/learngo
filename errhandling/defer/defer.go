package main

import (
	"fmt"
	"os"
	"bufio"
	"wuzhengyu/learngo/functional/fib"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile (filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib.Fibonacci()
	for i := 0; i< 20; i++ {
		fmt.Fprintln(write, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}

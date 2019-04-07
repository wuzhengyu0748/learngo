package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1101
		convertToBin(232478),
	)

	printFile("abc.txt")

	s := `abc"dd"
	llll
	kkkkwe
	
	23123'88'`
	printFileContents(strings.NewReader(s))

	//forever()
}

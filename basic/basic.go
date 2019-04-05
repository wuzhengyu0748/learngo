package basic

import (
	"fmt"
	"math"
)

var (
	aa = 3
	ss = "kkk"
 	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a,b,c,d = 1, 2, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 1, 2, true, "def"
	b = 5
	fmt.Println(a, b, c, d)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts() {
	const (
	 	filename string = "abc.txt";
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, python, golang, javascript)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
		eb
	)
	fmt.Println(b, kb, mb, gb, tb, pb, eb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	triangle()
	consts()
	enums()
}

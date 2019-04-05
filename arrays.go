package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100;
	for i,v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]bool
	fmt.Println(grid)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for _,v := range arr3 {
		fmt.Println(v)
	}

	printArray(&arr1)
	printArray(&arr3)
	fmt.Println(arr1, arr3)
}

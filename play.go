package main

import (
	"fmt"
)

func main() {

	arr := make([]int, 0, 10)
	fmt.Println(arr, len(arr), cap(arr))

	for i := 0; i < 1025; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr, len(arr), cap(arr))
}

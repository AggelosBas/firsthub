package student

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}


package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg1 := 4
	fmt.Println(piscine.Fibonacci(arg1))
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, err1 := strconv.Atoi(os.Args[1])
	op := os.Args[2]
	b, err2 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil {
		return
	}

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}
		fmt.Println(a / b)
	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		fmt.Println(a % b)
	default:
		return
	}
}

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	// Παίρνουμε τα arguments
	aStr := os.Args[1]
	operator := os.Args[2]
	bStr := os.Args[3]

	// Μετατρέπουμε σε int64
	a, err1 := strconv.ParseInt(aStr, 10, 64)
	b, err2 := strconv.ParseInt(bStr, 10, 64)

	// Αν κάποιο conversion απέτυχε → δεν εκτυπώνουμε τίποτα
	if err1 != nil || err2 != nil {
		return
	}

	switch operator {
	case "+":
		result := a + b
		// Έλεγχος overflow
		if (a > 0 && b > 0 && result < 0) || (a < 0 && b < 0 && result > 0) {
			return
		}
		fmt.Println(result)
	case "-":
		result := a - b
		if (b > 0 && a < math.MinInt64+b) || (b < 0 && a > math.MaxInt64+b) {
			return
		}
		fmt.Println(result)
	case "*":
		result := a * b
		// Overflow detection
		if b != 0 && result/b != a {
			return
		}
		fmt.Println(result)
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
		// Αν ο τελεστής δεν είναι έγκυρος → τίποτα
		return
	}
}

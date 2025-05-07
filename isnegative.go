package main

import "piscine-go/isnegative"

func IsNegative(n int) {
	if n < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

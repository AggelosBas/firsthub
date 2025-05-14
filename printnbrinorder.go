package student

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var count [10]int
	for g := n; g > 0; g /= 10 {
		digit := g % 10
		count[digit]++
	}
	for d := 0; d <= 9; d++ {
		for i := 0; i < count[d]; i++ {
			z01.PrintRune(rune('0' + d))
		}
	}
}

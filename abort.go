package piscine

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers[2]
}

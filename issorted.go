package student

func IsSorted(f func(a, b int) int, a []int) bool {
	isAscending := true
	isDescending := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			isAscending = false
		}
		if f(a[i], a[i+1]) < 0 {
			isDescending = false
		}
	}

	return isAscending || isDescending
}

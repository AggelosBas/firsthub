package student

func ShoppingListSort(slice []string) []string {
	n := len(slice)

	// Bubble Sort based on length
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				// Swap
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

package student

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	var result [][]string

	for i := 0; i < n; i++ {
		result = append(result, []string{})
	}

	for i := 0; i < n; i++ {
		result[i] = podium[n-1-i]
	}

	return result
}

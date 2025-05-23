package student

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	result := makeFixedSlice(n)

	for i := 0; i < n; i++ {
		result[i] = podium[n-1-i]
	}

	return result
}

func makeFixedSlice(n int) [][]string {
	var res [][]string
	for i := 0; i < n; i++ {
		res = append(res, []string{})
	}
	return res
}

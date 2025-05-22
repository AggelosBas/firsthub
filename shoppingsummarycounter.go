package student

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(str)

	for _, word := range words {
		result[word]++
	}

	return result
}

package student

func ShoppingSummaryCounter(str string) map[string]int {
	counts := Split(str, " ")
	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}
	return counts
}
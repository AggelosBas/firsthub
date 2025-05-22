package student

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		} else if word != "" {
			result[word]++
			word = ""
		}
	}
	if word != "" {
		result[word]++
	}
	return result
}

package student

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		} else {
			if word != "" {
				summary[word]++
				word = ""
			}
		}
	}
		summary[word]++
	return summary
}

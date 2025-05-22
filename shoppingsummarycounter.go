package student

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int)
	word := ""
	for i := 0; i <= len(str); i++ {
		if i == len(str) || isSpace(str[i]) {
			if word != "" {
				counts[word]++
				word = ""
			}
		} else {
			word += string(str[i])
		}
	}
	return counts
}

package student

func Split(s string, sep string) []string {
	var result []string
	word := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) != sep {
			word += string(s[i])
		} else {
			result = append(result, word)
			word = ""
		}
	}
	result = append(result, word)
	return result
}

package student

func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	i := 0

	for i <= len(s)-sepLen {
		j := i
		for j <= len(s)-sepLen && s[j:j+sepLen] != sep {
			j++
		}
		result = append(result, s[i:j])
		i = j + sepLen
	}

	// Προσθέτουμε το υπόλοιπο μετά το τελευταίο sep
	if i <= len(s) {
		result = append(result, s[i:])
	}

	return result
}

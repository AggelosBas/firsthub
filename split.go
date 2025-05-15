package student

func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	i := 0

	for i < len(s) {
		j := i
		// Προχωράμε μέχρι να βρούμε το sep ή να φτάσουμε στο τέλος
		for j <= len(s)-sepLen && s[j:j+sepLen] != sep {
			j++
		}

		// Προσθέτουμε τη λέξη που βρέθηκε
		result = append(result, s[i:j])

		// Αν φτάσαμε στο τέλος, σταματάμε
		if j > len(s)-sepLen {
			break
		}

		// Προχωράμε μετά το sep
		i = j + sepLen
	}

	return result
}

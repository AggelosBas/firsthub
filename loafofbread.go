package student

func LoafOfBread(str string) string {
	// Αφαιρούμε τα spaces
	cleaned := ""
	for _, r := range str {
		if r != ' ' {
			cleaned += string(r)
		}
	}

	if len(cleaned) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0
	for i+5 <= len(cleaned) {
		// Παίρνουμε 5 χαρακτήρες
		result += cleaned[i:i+5] + " "
		i += 6 // Προχωράμε 6 θέσεις (skip μετά το 5ο γράμμα)
	}

	// Αφαιρούμε το τελευταίο space και προσθέτουμε newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}

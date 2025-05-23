package student

func LoafOfBread(str string) string {
	// Βήμα 1: Αφαίρεση των spaces
	clean := ""
	for _, r := range str {
		if r != ' ' {
			clean += string(r)
		}
	}

	// Αν είναι κάτω από 5 χαρακτήρες, return
	if len(clean) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0

	for i+5 <= len(clean) {
		// Παίρνουμε 5 χαρακτήρες
		result += clean[i:i+5] + " "
		i += 6 // Παραλείπουμε τον 6ο χαρακτήρα
	}

	// Αφαιρούμε το τελευταίο space (αν υπάρχει) και προσθέτουμε newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}

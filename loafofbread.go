package student

func LoafOfBread(str string) string {
	chars := []rune{}
	// Αφαιρούμε τα spaces αλλά κρατάμε τα υπόλοιπα
	for _, r := range str {
		if r != ' ' {
			chars = append(chars, r)
		}
	}

	if len(chars) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0
	for i+5 <= len(chars) {
		// Πάρε 5 χαρακτήρες
		for j := 0; j < 5; j++ {
			result += string(chars[i+j])
		}
		result += " "
		i += 6 // Προχώρα 5 και μετά +1 (skip)
	}

	// Αφαίρεσε το τελευταίο space και βάλε newline
	if len(result) > 0 {
		result = result[:len(result)-1] + "\n"
	} else {
		result = "Invalid Output\n"
	}
	return result
}

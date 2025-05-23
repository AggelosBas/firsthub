package student

func LoafOfBread(str string) string {
	// Καθαρίζουμε τα spaces
	clean := ""
	for _, r := range str {
		if r != ' ' {
			clean += string(r)
		}
	}

	if len(clean) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	for i := 0; i+5 <= len(clean); i += 6 {
		result += clean[i : i+5]
		// Πρόσθεσε space αν υπάρχουν άλλα 5 χαρακτήρες μετά
		if i+6+4 < len(clean) {
			result += " "
		}
	}
	result += "\n"
	return result
}

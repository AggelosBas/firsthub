package student

func LoafOfBread(str string) string {
	count := 0
	word := ""
	result := ""
	skipping := false

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skipping {
			skipping = false
			continue
		}

		word += string(r)
		count++

		if count == 5 {
			result += word + " "
			word = ""
			count = 0
			skipping = true // Skip next character
		}
	}

	// Αν δεν δημιουργήθηκε ούτε μία λέξη
	if len(result) == 0 {
		return "Invalid Output\n"
	}

	// Αφαίρεσε τελευταίο κενό και πρόσθεσε newline
	return result[:len(result)-1] + "\n"
}

package student

func LoafOfBread(str string) string {
	// Βήμα 1: φτιάξε string χωρίς κενά
	cleaned := ""
	for _, r := range str {
		if r != ' ' {
			cleaned += string(r)
		}
	}

	// Αν έχει λιγότερους από 5 χαρακτήρες, return
	if len(cleaned) < 5 {
		return "Invalid Output\n"
	}

	// Βήμα 2: πάρε ομάδες των 5 και skip 1
	result := ""
	i := 0
	for i+5 <= len(cleaned) {
		result += cleaned[i:i+5] + " "
		i += 6 // 5 + skip 1
	}

	// Βήμα 3: αφαίρεσε το τελευταίο κενό και βάλε \n
	result = result[:len(result)-1] + "\n"
	return result
}

package student

func LoafOfBread(str string) string {
	// Αφαίρεση των κενών
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
		result += cleaned[i:i+5] + " "
		i += 6 // μετά από 5 γράμματα, skip 1
	}

	// Αν περισσεύει κάτι στο τέλος που δεν έγινε πεντάδα
	if i-1 < len(cleaned) {
		// Προσθέτουμε ό,τι μένει (όχι άλλο skip)
		remaining := cleaned[i-1:]
		result += remaining
	}

	// Αφαίρεση τελευταίου κενού αν υπάρχει
	if result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}

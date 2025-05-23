package student

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0
	word := ""

	for _, r := range str {
		if r == ' ' {
			continue
		}
		word += string(r)
		count++

		if count == 5 {
			result += word + " "
			word = ""
			count = 0
			// skip 1 character (space or not), if exists
			if len(str) > 0 {
				str = str[1:]
			}
		}
	}

	if len(result) == 0 {
		return "Invalid Output\n"
	}

	// Καθαρίζουμε το τελικό space και προσθέτουμε newline
	if result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result + "\n"
}

package student

func Split(s, sep string) []string {
	var result []string // Αποθήκευση αποτελεσμάτων
	sepLen := len(sep)  // Μήκος separator π.χ. "HA" = 2
	start := 0          // Αρχική θέση για κάθε κομμάτι

	for i := 0; i <= len(s)-sepLen; {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i]) // Προσθέτουμε ό,τι βρήκαμε μέχρι τώρα
			i += sepLen                         // Προχωράμε μετά τον separator
			start = i                           // Νέα αρχή
		} else {
			i++
		}
	}

	// Τελευταίο κομμάτι που απομένει (ό,τι είναι μετά το τελευταίο separator)
	if start <= len(s) {
		result = append(result, s[start:])
	}

	return result
}

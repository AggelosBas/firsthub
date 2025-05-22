package student

func Rot14(s string) string {
	var result string

	for i := 0; i < len(s); i++ {
		r := s[i]
		if r >= 'a' && r <= 'z' {
			newR := r + 14
			if newR > 'z' {
				newR = newR - 26
			}
			result += string(newR)
		} else if r >= 'A' && r <= 'Z' {
			newR := r + 14
			if newR > 'Z' {
				newR = newR - 26
			}
			result += string(newR)
		} else {
			result += string(r)
		}
	}

	return result
}

package student

func TrimAtoi(s string) int {
	sign := 1
	foundDigit := false
	n := 0
	for _, c := range s {

		if c == '-' && !foundDigit {
			sign = -1
		}
		if c >= '0' && c <= '9' {
			foundDigit = true
			n = n*10 + int(c-'0')
		}
	}
	if !foundDigit {
		return 0
	}
	return sign * n
}

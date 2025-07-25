package student

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= '0' && s[i] <= '9') ||
			(s[i] >= 'A' && s[i] <= 'Z') ||
			(s[i] >= 'a' && s[i] <= 'z')) ||
			s[i] == ' ' {
			return false
		}
	}
	return true
}

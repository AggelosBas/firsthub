package student

func ToLower(s string) string {
	b := []byte(s)
	for i, s := range b {
		if s >= 'A' && s <= 'Z' {
			b[i] = s + 32
		}
	}
	return string(b)
}

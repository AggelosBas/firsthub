package student

func ToUpper(s string) string {
	b := []byte(s)
	for i, s := range b {
		if s >= 'a' && s <= 'z' {
			b[i] = s - 32
		}
	}
	return string(b)
}

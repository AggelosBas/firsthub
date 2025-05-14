package student

func Capitalize(s string) string {
	result := []rune(s)
	isletter := false

	for i, c := range result {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			if !isletter {
				if c >= 'a' && c <= 'z' {
					result[i] = c - ('a' - 'A')
				}
				isletter = true
			} else {
				if c >= 'A' && c <= 'Z' {
					result[i] = c + ('a' - 'A')
				}
			}
		} else {
			isletter = false
		}
	}
	return string(result)
}

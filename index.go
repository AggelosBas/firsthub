package student

func Index(s string, toFind string) int {
	a := len(toFind)
	if a == 0 {
		return 0
	}
	for i := 0; i <= len(s)-a; i++ {
		if s[i:i+a] == toFind {
			return i
		}
	}
	return -1
}

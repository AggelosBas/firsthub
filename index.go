package student

func Index(s string, toFind string) int {
	sLen := len(toFind)
	if a == 0{
		return 0
	}
		for i := 0; i <= sLen-toFindLen; i++ {
		if s[i:i+toFindLen] == toFind {
			return i
		}
	}
	return -1
}

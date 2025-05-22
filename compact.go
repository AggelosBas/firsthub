package student

func Compact(ptr *[]string) int {
	a := *ptr
	j := 0
	for i := 0; i < len(a); i++ {
		if a[i] != "" {
			a[j] = a[i]
			j++
		}
	}
	*ptr = a[:j]
	return j
}

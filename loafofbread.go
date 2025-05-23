package student

func LoafOfBread(str string) string {
	s := []rune(str)
	var rs []rune
	if len(s) == 0 {
		return "\n"
	}
	if len(s) < 5 {
		return "Invalid input\n"
	}

	counter := 0
	for _, c := range s {
		if counter == 5 {
			rs = append(rs, ' ')
			counter = 0
			continue
		}
		if c == ' ' {
			continue
		}
		rs = append(rs, c)
		counter++
	}

	for len(rs) > 0 && rs[len(rs)-1] == ' ' {
		rs = rs[:len(rs)-1]
	}

	rs = append(rs, '\n')
	return string(rs)
}

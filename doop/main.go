package main

import (
	"os"
)

func Atoi(s string) (int64, bool) {
	var res int64 = 0
	sign := int64(1)
	i := 0

	if len(s) == 0 {
		return 0, false
	}

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		res = res*10 + int64(s[i]-'0')
	}

	return res * sign, true
}

func Itoa(n int64) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	var res [20]byte
	i := len(res)

	for n > 0 {
		i--
		res[i] = byte(n%10) + '0'
		n /= 10
	}

	return sign + string(res[i:])
}

func Println(s string) {
	os.Stdout.Write([]byte(s))
	os.Stdout.Write([]byte{'\n'})
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, ok1 := Atoi(aStr)
	b, ok2 := Atoi(bStr)

	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		res := a + b
		// Προσπαθούμε απλή προστασία από overflow
		if (a > 0 && b > 0 && res < 0) || (a < 0 && b < 0 && res > 0) {
			return
		}
		Println(Itoa(res))
	case "-":
		res := a - b
		if (a > 0 && b < 0 && res < 0) || (a < 0 && b > 0 && res > 0) {
			return
		}
		Println(Itoa(res))
	case "*":
		res := a * b
		if b != 0 && res/b != a {
			return
		}
		Println(Itoa(res))
	case "/":
		if b == 0 {
			Println("No division by 0")
			return
		}
		Println(Itoa(a / b))
	case "%":
		if b == 0 {
			Println("No modulo by 0")
			return
		}
		Println(Itoa(a % b))
	default:
		return
	}
}

package student

import (
	"math"
)

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	sqrt := math.Sqrt(float64(nb))
	if sqrt == math.Floor(sqrt) {
		return int(sqrt)
	}
	return 0
}

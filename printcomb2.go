package piscine

import "fmt"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			fmt.Printf("%02d %02d", i, j)
			if i != 98 || j != 99 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()
}

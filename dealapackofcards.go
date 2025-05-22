package student

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < 3; j++ {
			cardIndex := i*3 + j
			if j < 2 {
				fmt.Printf("%d, ", deck[cardIndex])
			} else {
				fmt.Printf("%d", deck[cardIndex])
			}
		}
		fmt.Println()
	}
}

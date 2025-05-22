package main

import (
	"fmt"
	"student"
)

func main() {
	fmt.Println(student.FoodDeliveryTime("burger"))
	fmt.Println(student.FoodDeliveryTime("chips"))
	fmt.Println(student.FoodDeliveryTime("nuggets"))
	fmt.Println(student.FoodDeliveryTime("burger") +
		student.FoodDeliveryTime("chips") +
		student.FoodDeliveryTime("nuggets"))
}

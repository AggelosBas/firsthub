package main

import "fmt"

// ✅ Δηλώνουμε τον τύπο Pilot σύμφωνα με τις απαιτήσεις της main
type Pilot struct {
	Name     string
	Life     float64
	Age      int
	Aircraft int
}

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

const AIRCRAFT1 = 1

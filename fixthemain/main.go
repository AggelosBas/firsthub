package main

import "piscine"

func main() {
	door := &piscine.Door{}

	piscine.OpenDoor(door)
	if piscine.IsDoorClose(door) {
		piscine.OpenDoor(door)
	}
	if piscine.IsDoorOpen(*door) {
		piscine.CloseDoor(door)
	}
	if door.state == piscine.OPEN {
		piscine.CloseDoor(door)
	}
}

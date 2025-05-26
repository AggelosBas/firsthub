package student

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	index := 0

	for current != nil {
		if index == pos {
			return current
		}
		current = current.Next
		index++
	}
	return nil // Αν δεν βρεθεί ή ξεπεράσει τα όρια
}

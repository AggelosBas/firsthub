package student

func ListReverse(l *List) {
	var prev *NodeL = nil
	current := l.Head
	l.Tail = l.Head // η παλιά Head γίνεται Tail

	for current != nil {
		next := current.Next // κρατάμε τον επόμενο
		current.Next = prev  // αντιστρέφουμε τον δείκτη
		prev = current       // ενημερώνουμε το προηγούμενο
		current = next       // πάμε στον επόμενο
	}

	l.Head = prev // η τελευταία τιμή του prev είναι η νέα Head
}

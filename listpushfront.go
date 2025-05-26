package student

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil { // Αν η λίστα είναι άδεια
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head // Το νέο node δείχνει στο παλιό Head
		l.Head = newNode      // Το νέο node γίνεται το νέο Head
	}
}

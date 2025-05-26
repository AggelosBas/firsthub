package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil { // λίστα άδεια
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode // το παλιό tail δείχνει στο νέο node
		l.Tail = newNode      // το νέο node γίνεται το νέο tail
	}
}

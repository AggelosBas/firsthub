package student

func SortListInsert(l *NodeI, data_ref int) *NodeI {
    newNode := &NodeI{Data: data_ref}

    // Περίπτωση: λίστα είναι άδεια ή το νέο στοιχείο είναι μικρότερο από την κεφαλή
    if l == nil || data_ref < l.Data {
        newNode.Next = l
        return newNode
    }

    current := l
    // Βρίσκουμε τη θέση για εισαγωγή
    for current.Next != nil && current.Next.Data < data_ref {
        current = current.Next
    }

    // Εισάγουμε τον νέο κόμβο
    newNode.Next = current.Next
    current.Next = newNode

    return l
}

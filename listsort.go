package student

type NodeI struct {
    Data int
    Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
    if l == nil || l.Next == nil {
        return l
    }

    mid := getMiddle(l)
    right := mid.Next
    mid.Next = nil

    leftSorted := ListSort(l)
    rightSorted := ListSort(right)

    return merge(leftSorted, rightSorted)
}

func getMiddle(l *NodeI) *NodeI {
    slow, fast := l, l
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

func merge(l1, l2 *NodeI) *NodeI {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    if l1.Data < l2.Data {
        l1.Next = merge(l1.Next, l2)
        return l1
    } else {
        l2.Next = merge(l1, l2.Next)
        return l2
    }
}

package student

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// Αν ο node είναι η ρίζα
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}

package student

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		// Δημιουργούμε τον πρώτο κόμβο αν δεν υπάρχει δέντρο
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		// Πήγαινε αριστερά
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		// Πήγαινε δεξιά (αν είναι ίσο ή μεγαλύτερο)
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}

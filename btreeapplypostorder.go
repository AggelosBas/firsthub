package student

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyPostorder(root.Left, f)   // 1. Επισκέπτεται το αριστερό παιδί
	BTreeApplyPostorder(root.Right, f)  // 2. Επισκέπτεται το δεξί παιδί
	f(root.Data) // 3. Εφαρμόζει τη συνάρτηση στον τρέχοντα κόμβο
}

package functions

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(1+maxDepth(root.Right)+maxDepth(root.Left), diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

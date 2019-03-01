package easy

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l
	} else {
		return r
	}
}

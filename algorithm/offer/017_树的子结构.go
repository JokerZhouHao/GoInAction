package main

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func isSubtree(root1, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	} else if root1 == nil {
		return false
	}
	if root1.val == root2.val && isSubtree(root1.left, root2.left) && isSubtree(root1.right, root2.right) {
		return true
	}
	return false
}

func HasSubtree(root1, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.val == root2.val && isSubtree(root1.left, root2.left) && isSubtree(root1.right, root2.right) {
		return true
	}
	return HasSubtree(root1.left, root2) || HasSubtree(root1.right, root2)

}

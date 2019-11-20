package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func index(arr []int, start, end, x int) int {
	for i := start; i <= end; i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}

func r(pre []int, ps, pe int, in []int, is, ie int) *TreeNode {
	if pe < ps {
		return nil
	} else if pe == ps {
		return &TreeNode{pre[ps], nil, nil}
	}

	mid := index(in, is, ie, pre[ps])
	node := TreeNode{pre[ps], nil, nil}
	node.left = r(in, ps+1, ps+mid-is, in, is, mid-1)
	node.right = r(in, ps+mid-is+1, pe, in, mid+1, ie)
	return &node
}

func reConstructBinaryTree(pre []int, in []int) *TreeNode {
	return r(pre, 0, len(pre), in, 0, len(in))
}

func main() {

}

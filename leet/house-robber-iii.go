package leet

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func RobTree(root *TreeNode) int {

	//recurance relation
	//max(root.Left,
	return maxArr(InOrder(root))
	//return 0
}

func InOrder(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l := InOrder(root.Left)
	r := InOrder(root.Right)
	withoutNode := maxArr(l) + maxArr(r)
	node := root.Val + l[1] + r[1]
	return []int{node, withoutNode}
}

func maxArr(a []int) int {
	if a[0] < a[1] {
		return a[1]
	}
	return a[0]
}

package leet

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Twosumivinputisabst(root *TreeNode, k int) bool {
	var a []int
	inorder(root, &a)
	l := 0
	r := len(a) - 1

	for l < r {
		sum := a[l] + a[r]
		if sum == k {
			return true
		} else if sum < k {
			l++
		} else {
			r--
		}
	}
	return false
}

func inorder(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, a)
	*a = append(*a, root.Val)
	inorder(root.Right, a)
}

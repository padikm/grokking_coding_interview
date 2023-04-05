package ds

type Bst struct {
	D     int
	Left  *Bst
	Right *Bst
}

func Insert(b *Bst, i int) {
	if b == nil {
		b = &Bst{D: i}
		return
	}
	if b.D > i {
		if b.Left == nil {
			b.Left = &Bst{D: i}
			return
		}
		Insert(b.Left, i)
	} else {
		if b.Right == nil {
			b.Right = &Bst{D: i}
			return
		}
		Insert(b.Right, i)
	}

}

func InorderSuccessor(b *Bst) *Bst {
	tmp := b
	tmp = tmp.Right
	for tmp.Left != nil {
		tmp = tmp.Left
	}
	return tmp
}

func DeleteBstNode(b *Bst, i int) {
	DeleteBstANode(b, nil, i)
}

func DeleteBstANode(b *Bst, prev *Bst, i int) {
	if b == nil {
		return
	}
	if i < b.D {
		DeleteBstANode(b.Left, b, i)
	} else if i > b.D {
		DeleteBstANode(b.Right, b, i)
	} else {
		if b.Left == nil && b.Right == nil {
			if prev.Left != nil && prev.Left.D == i {
				prev.Left = nil
			} else if prev.Right != nil && prev.Right.D == i {
				prev.Right = nil
			}
			return
		} else if b.Left == nil {
			b = b.Right
			return
		} else if b.Right == nil {
			b = b.Left
		}
		temp := InorderSuccessor(b)
		b.D = temp.D
		DeleteBstANode(b.Right, b, temp.D)
	}
}

func SearchNode(b *Bst, i int) bool {
	if b == nil {
		return false
	}
	if b.D == i {
		return true
	}
	if b.D > i {
		return SearchNode(b.Left, i)
	} else {
		return SearchNode(b.Right, i)
	}
}

//func insertIntoBST(root *TreeNode, val int) *TreeNode {
//	if root == nil {
//		return &TreeNode{Val:val}
//	}
//
//	if root.Val > val {
//		root.Left = insertIntoBST(root.Left,val)
//	} else {
//		root.Right = insertIntoBST(root.Right,val)
//	}
//
//	return root
//}

func ValidateBst(b *Bst) bool {
	if b == nil {
		return true
	}

	if b.Left != nil && b.D < b.Left.D {
		return false
	}
	if b.Right != nil && b.D > b.Right.D {
		return false
	}

	if !ValidateBst(b.Left) || !ValidateBst(b.Right) {
		return false
	}
	return true
}

func IsBST(b, l, r *Bst) bool {
	if b == nil {
		return true
	}

	if l != nil && b.D <= l.D {
		return false
	}
	if r != nil && b.D >= r.D {
		return false
	}

	return IsBST(b.Left, l, b) && IsBST(b.Right, b, r)
}

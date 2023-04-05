package ds

import (
	"fmt"
)

type Itree interface {
	Insert()
	DeleteLeaf()
}

type Tree struct {
	d     int
	left  *Tree
	right *Tree
}

func (t *Tree) Insert() {
	t1 := &Tree{d: 1}
	t2 := &Tree{d: 2}
	t3 := &Tree{d: 3}
	t4 := &Tree{d: 4}
	t5 := &Tree{d: 5}
	t6 := &Tree{d: 6}
	t.left = t1
	t.right = t2
	t.left.left = t3
	t.left.right = t4
	t.right.left = t5
	t.right.right = t6
	t.right.right.left = &Tree{d: 7}
}

func (t *Tree) DeleteLeaf() {

}

func Inorder(t *Tree) {
	if t == nil {
		return
	}
	Inorder(t.left)

	fmt.Println(t.d)
	Inorder(t.right)
}

func Preorder(t *Tree) {
	if t == nil {
		return
	}
	fmt.Println(t.d)
	Preorder(t.left)
	Preorder(t.right)
}

func Postorder(t *Tree) {
	if t == nil {
		return
	}
	Postorder(t.left)
	Postorder(t.right)
	fmt.Println(t.d)
}

func DepthOfTree(t *Tree) int {
	if t == nil {
		return 0
	}
	ld := DepthOfTree(t.left)
	rd := DepthOfTree(t.right)
	return max(ld, rd) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func InorderBst(b *Bst) {
	if b == nil {
		return
	}
	InorderBst(b.Left)
	fmt.Println(b.D)
	InorderBst(b.Right)
}

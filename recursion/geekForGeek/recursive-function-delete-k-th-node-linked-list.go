package geekForGeek

import (
	"awesomeProject/educative"
	"fmt"
)

func RecursiveFunctionDeleteKthNodeLinkedList(node *educative.ListNode, k int) {
	node = recursiveFunctionDeleteKthNodeLinkedList(node, k)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func recursiveFunctionDeleteKthNodeLinkedList(node *educative.ListNode, k int) *educative.ListNode {
	if k == 0 {
		return node
	}
	if k == 1 {
		return node.Next
	}

	node.Next = recursiveFunctionDeleteKthNodeLinkedList(node.Next, k-1)
	//if k == 1 {
	//	node.Next = res.Next
	//}

	return node

}

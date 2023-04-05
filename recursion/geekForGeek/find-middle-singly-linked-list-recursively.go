package geekForGeek

import (
	"awesomeProject/educative"
	"fmt"
)

func FindMiddleOfListRec(node *educative.ListNode) {
	res := findMiddleOfListRec(node, node)
	fmt.Println(res.Val)
}

func findMiddleOfListRec(node *educative.ListNode, node1 *educative.ListNode) *educative.ListNode {

	if node == nil || node1 == nil || node1.Next == nil {
		return node
	}

	return findMiddleOfListRec(node.Next, node1.Next.Next)

}

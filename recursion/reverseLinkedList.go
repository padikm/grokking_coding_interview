package recursion

import (
	"awesomeProject/educative"
)

func ReverseLinkedList(head *educative.ListNode) *educative.ListNode {
	cur := head
	var prev *educative.ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func ReverseListRecursion(head *educative.ListNode) *educative.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := ReverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

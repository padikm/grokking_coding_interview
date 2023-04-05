package educative

import "fmt"

func ReorderList(head *ListNode) {
	//find middle
	cur := head
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	midRev := revFromMid(slow)
	fmt.Println(midRev)
	for cur != nil && midRev != nil {
		next := cur.Next
		cur.Next = midRev
		midRevNext := midRev.Next
		midRev.Next = next
		cur = next
		midRev = midRevNext
	}
	if cur != nil {
		cur.Next = nil
	}
	if midRev != nil {
		midRev.Next = nil
	}
	fmt.Println(head)
}

func revFromMid(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

package educative

func MiddleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}

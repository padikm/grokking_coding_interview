package educative

func IsPalindrome(head *ListNode) bool {
	f := head
	s := head
	var prev *ListNode
	for f != nil && f.Next != nil {
		f = f.Next.Next
		prev = s
		s = s.Next
	}
	if prev == nil {
		return true
	}
	prev.Next = revList(s)
	f = head
	s = head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		prev = s
		s = s.Next
	}
	mid := s
	tem := head
	for tem != mid {
		if tem.Val != s.Val {
			return false
		}
		tem = tem.Next
		s = s.Next
	}
	return true
}

func revList(head *ListNode) *ListNode {
	cur := head
	var next, prev *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

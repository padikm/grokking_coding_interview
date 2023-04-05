package educative

func SwapPairs(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil && cur.Next != nil {
		prevTmp, curTmp := revSwapPair(cur, 2)
		if prev == nil {
			head = prevTmp
			prev = cur
		} else {
			prev.Next = prevTmp
			prev = cur
		}
		cur.Next = curTmp
		cur = curTmp
	}
	return head
}

func revSwapPair(cur *ListNode, l int) (*ListNode, *ListNode) {
	var prev *ListNode
	for l != 0 {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		l--
	}
	return prev, cur
}

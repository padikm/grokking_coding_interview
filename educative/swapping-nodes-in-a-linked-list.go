package educative

func SwapNodes(head *ListNode, k int) *ListNode {
	cur := head
	//count nodes
	count := 0
	countNodeSwap := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	// find node to be swapped
	countNodeSwap = count - k
	cur = head
	for countNodeSwap != 0 {
		cur = cur.Next
		countNodeSwap--
	}
	lastCur := cur
	cur = head
	for k != 1 {
		cur = cur.Next
		k--
	}
	tmp := cur.Val
	cur.Val = lastCur.Val
	lastCur.Val = tmp

	return head

}

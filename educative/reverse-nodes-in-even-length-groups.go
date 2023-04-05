package educative

func ReverseEvenLengthGroups(head *ListNode) *ListNode {
	count := 0
	cur := head
	var prev *ListNode
	for cur != nil {
		count++
		cur = cur.Next
	}
	if count == 1 {
		return head
	}
	cur = head
	totalProcessed := 0
	for i := 1; cur != nil; i++ {
		rem := count - totalProcessed
		if (i%2 == 0 && (i <= rem)) || (i > rem && rem%2 == 0) {
			prevTmp, curNew := revGivenGroup(cur, i)
			prev.Next = prevTmp
			cur.Next = curNew
			prev = cur
			cur = curNew
			totalProcessed += i
		} else {
			for j := 0; j < i; j++ {
				if cur == nil {
					break
				}
				prev = cur
				cur = cur.Next
				totalProcessed++
			}
		}
	}
	return head
}

func revGivenGroup(head *ListNode, l int) (*ListNode, *ListNode) {
	cur := head
	var prev *ListNode
	for cur != nil && l != 0 {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		l--
	}
	return prev, cur

}

package educative

func ReverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	var prev, lastPrev, lastCur, tmp *ListNode
	//var next *ListNode
	//count nodes
	count := 0
	cNodes := countNodes(head)
	if k == 1 || head == nil || cNodes < k {
		return head
	}
	for cur != nil {
		if count == 1 {
			tmp = prev
			lastCur = prev
		}
		if count == k {
			if lastPrev != nil {
				lastPrev.Next = prev
				lastPrev = lastCur
			} else {
				head = prev
				lastPrev = tmp
			}

			count = 1
			if countNodes(cur) < k {
				lastCur.Next = cur
				break
			}
			lastCur = cur
			prev = cur
			cur = cur.Next
		}
		count++
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next

	}
	if cNodes == k {
		return prev
	}
	if cur == nil {
		lastPrev.Next = prev
		lastCur.Next = cur
	}
	return head
}

func countNodes(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

package educative

func ReverseBetween(head *ListNode, left int, right int) *ListNode {

	cur := head
	count := 1
	var leftPrevCur, prev *ListNode
	for count != left {
		count++
		leftPrevCur = cur
		cur = cur.Next
	}
	//rev from cur to right
	leftNode := cur
	for cur != nil && count <= right {
		count++
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	leftNode.Next = cur
	if leftPrevCur != nil {
		leftPrevCur.Next = prev
	} else {
		return prev
	}

	//fix the rev part
	return head

}

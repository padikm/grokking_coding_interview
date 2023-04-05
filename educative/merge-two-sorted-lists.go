package educative

func MergeTwoListsItr(list1 *ListNode, list2 *ListNode) *ListNode {
	var tmp, res *ListNode
	if list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp = list1
			list1 = list1.Next
		} else {
			tmp = list2
			list2 = list2.Next
		}
	}
	if tmp != nil {
		res = tmp
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			tmp = tmp.Next
			list1 = list1.Next
		} else {
			tmp.Next = list2
			tmp = tmp.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		if tmp == nil {
			tmp = list1
			res = tmp
		} else {
			tmp.Next = list1
		}
	}

	if list2 != nil {
		if tmp == nil {
			tmp = list2
			res = tmp
		} else {
			tmp.Next = list2
		}
	}
	return res
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}
}

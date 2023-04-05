package leet

import (
	"sort"
)

/**
 * Definition for singly-linked list.
 * type ListNode structtest {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	temp := head
	arr := make([]int, 0)
	for temp != nil {
		arr = append(arr, temp.Val)
		temp = temp.Next
	}
	sort.Ints(arr)
	temp = head
	i := 0
	for temp != nil && i < len(arr) {
		temp.Val = arr[i]
		i++
		temp = temp.Next
	}
	return head
}

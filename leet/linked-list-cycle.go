package leet

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode structtest {
 *     Val int
 *     Next *ListNode
 * }
 */

/*

l := &leet.ListNode{}
	a := []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}
	pos := 24
	head := l
	count := 0
	var saved *leet.ListNode
	for i, v := range a {
		l.Val = v
		l.Next = &leet.ListNode{}
		if i == pos {
			saved = l
		}
		count++

		if count == len(a) {
			l.Next = saved
		}
		l = l.Next
	}

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	var fp, sp *ListNode
	if head != nil && head.Next != nil {
		fp = head
		sp = head.Next.Next
	}
	for fp != nil && sp != nil && sp.Next != nil {
		if fp == sp {
			return true
		}
		fp = fp.Next
		sp = sp.Next.Next
	}
	return false
}

func DetectCycle(head *ListNode) *ListNode {
	var fp, sp *ListNode
	if head != nil && head.Next != nil {
		fp = head.Next
		sp = head.Next.Next
	}
	flag := false
	count := 0
	for fp != nil && sp != nil && sp.Next != nil {
		if fp == sp {
			//fmt.Println(sp.Val)
			flag = true
			break
		}
		fp = fp.Next
		sp = sp.Next.Next

	}
	fp = head
	if flag {
		for fp != sp {
			count++
			fp = fp.Next
			sp = sp.Next
		}
		fmt.Println("Count ", count)
		return fp
	}
	return nil
}

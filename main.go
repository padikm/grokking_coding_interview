package main

import (
	"awesomeProject/educative"
	"awesomeProject/leet"
	"fmt"
)

func main() {
	l := &educative.ListNode{}
	p := []int{1, 2, 3, 4, 5, 6}
	pos := -1
	//head := l
	count := 0
	var saved *educative.ListNode
	for i, v := range p {
		l.Val = v
		l.Next = &educative.ListNode{}
		if i == pos {
			saved = l
		}
		count++

		if count == len(p) {
			l.Next = saved
		}
		l = l.Next
	}
	/*
	   10
	   0
	   {1,2,3}
	   {0,1,2}
	*/
	//c := []int{0, 0, 0}
	//nums2 := []int{2, 4, 6}
	//mat := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	//fmt.Println(recursion.SkipCharFromStr("abbsbbaabsfgr", "a"))

	//recursion.NormalTriangle(0, 0, 5)
	//recursion.PrintSubSeqOfStr("inclu")
	fmt.Println(leet.TwoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}))
}

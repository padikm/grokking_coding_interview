package main

import (
	"awesomeProject/backtracking/leet"
	"awesomeProject/educative"
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
	//fmt.Println(leet.RobCircle([]int{1, 2, 1, 1}))
	//s := [][]string{{"5", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"}}
	//s := [][]string{{".", ".", ".", ".", ".", ".", ".", ".", "."}, {"4", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", "6", ".", "."}, {".", ".", ".", "3", "8", ".", ".", ".", "."}, {".", "5", ".", ".", ".", "6", ".", ".", "1"}, {"8", ".", ".", ".", ".", ".", ".", "6", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", "7", ".", "9", ".", ".", ".", "."}, {".", ".", ".", "6", ".", ".", ".", ".", "."}}

	fmt.Println(leet.LetterCasePermutation("abcd"))
}

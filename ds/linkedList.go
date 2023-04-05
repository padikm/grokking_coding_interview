package ds

import (
	"fmt"
)

type ILinkedList interface {
	Insert(int)
	DeleteLast()
	InsertFront(int)
	Display()
	Reverse() *LinkedList
	DeleteGivenNode(int)
	DetectLoop() int
}

type LinkedList struct {
	d    int
	link *LinkedList
}

func (head *LinkedList) Insert(i int) {
	node := &LinkedList{d: i}
	if head == nil {
		head = node
	}
	temp := head
	for temp.link != nil {
		temp = temp.link
	}
	temp.link = node
}

func (head *LinkedList) DeleteLast() {
	if head.link == nil {
		return
	}
	temp := head.link
	prev := head.link
	for temp.link != nil {
		prev = temp
		temp = temp.link
	}

	prev.link = nil
}

func (head *LinkedList) Display() {
	if head.link == nil {
		return
	}
	temp := head.link
	for temp.link != nil {
		fmt.Println(temp.d)
		temp = temp.link
	}
	fmt.Println(temp.d)
}

func (head *LinkedList) Reverse() *LinkedList {
	var prev *LinkedList
	var next *LinkedList
	head = head.link
	for head != nil {
		next = head.link
		head.link = prev
		prev = head
		head = next
	}
	return prev
}

func (head *LinkedList) InsertFront(i int) {
	node := &LinkedList{d: i}
	if head == nil {
		head = node
	}
	next := head.link
	head.link = node
	node.link = next
}

func (head *LinkedList) DeleteGivenNode(i int) {
	if head == nil || head.link == nil {
		return
	}
	next := head.link
	prev := head
	for next != nil {
		if next.d == i {
			break
		}
		prev = next
		next = next.link
	}
	if next != nil {
		prev.link = next.link
	}
}

func (head *LinkedList) DetectLoop() int {
	if head == nil || head.link == nil {
		return -1
	}
	//create loop
	func(head *LinkedList) {
		if head == nil || head.link == nil {
			return
		}
		//change the next to create loop
		next := head.link
		nextToNext := head.link
		for nextToNext.link != nil {
			nextToNext = nextToNext.link
		}

		nextToNext.link = next

	}(head)
	next := head.link
	nextToNext := head.link.link
	var prev *LinkedList
	for next.d != nextToNext.d {
		prev = next
		next = next.link
		nextToNext = nextToNext.link.link
	}
	return prev.d
}

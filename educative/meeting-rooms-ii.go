package educative

import (
	"container/heap"
	"fmt"
)

type PriorityQueue [][]int

func (p *PriorityQueue) Len() int {
	return len(*p)
}

func (p *PriorityQueue) Less(i, j int) bool {
	return (*p)[i][1] < (*p)[j][1]
}

func (p *PriorityQueue) Swap(i, j int) {
	temp := (*p)[i][1]
	(*p)[i][1] = (*p)[j][1]
	(*p)[j][1] = temp
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.([]int))
}

func (p *PriorityQueue) Pop() any {
	old := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return old
}

func MeetingRoomsIi(a [][]int) [][]int {
	pq := PriorityQueue(a)
	heap.Init(&pq)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).([]int)
		fmt.Println(item)
	}
	return [][]int{}
}

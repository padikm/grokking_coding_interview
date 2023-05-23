package leet

import (
	"container/heap"
)

type maxHeap [][]int

func (m *maxHeap) Len() int {
	return len(*m)
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i][1] > (*m)[j][1]
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *maxHeap) Pop() interface{} {
	i := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return i
}

func MinRefuelStops(target int, startFuel int, stations [][]int) int {
	mh := maxHeap(stations)
	heap.Init(&mh)
	count := 0
	dist := startFuel
	//l := len(stations)
	tmpArr := make([][]int, 0)
	if startFuel >= target {
		return 0
	}

	for len(mh) > 0 {
		res := heap.Pop(&mh).([]int)
		if res[0] > dist {
			//heap.Push(&mh, res)
			tmpArr = append(tmpArr, res)
		} else {
			dist = dist + res[1]
			for _, a := range tmpArr {
				heap.Push(&mh, a)
			}
			tmpArr = make([][]int, 0)
			count++
		}
		if dist >= target {
			break
		}

	}

	if len(mh) == 0 && dist < target {
		return -1
	}
	return count
}

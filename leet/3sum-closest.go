package leet

import (
	"container/heap"
	"math"
	"sort"
)

//Heap not needed :(
type minDiffHeap [][]int

func (m *minDiffHeap) Len() int {
	//TODO implement me
	return len(*m)
}

func (m *minDiffHeap) Less(i, j int) bool {
	//TODO implement me
	return (*m)[i][0] < (*m)[j][0]
}

func (m *minDiffHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minDiffHeap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *minDiffHeap) Pop() interface{} {
	small := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return small
}

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minHeap := minDiffHeap{}
	heap.Init(&minHeap)
	for i := 0; i < len(nums)-2; i++ {
		threeSum(nums, &minHeap, nums[i], i+1, target)
	}
	diff := math.Abs(float64(target - minHeap[0][0]))
	sum := minHeap[0][1] + minHeap[0][2] + minHeap[0][3]
	for len(minHeap) > 0 {
		element := heap.Pop(&minHeap).([]int)
		curDiff := float64(target - element[0])
		tempSum := element[1] + element[2] + element[3]
		if math.Abs(curDiff) < diff {
			diff = curDiff
			sum = tempSum
		}
	}

	return sum
}

func threeSum(a []int, minHeap *minDiffHeap, firstNum, low, target int) {
	high := len(a) - 1
	//tmp := a[low] + a[high] + firstNum
	for low < high {
		rem := firstNum + a[low] + a[high]
		heap.Push(minHeap, []int{rem, firstNum, a[low], a[high]})
		if rem > target {
			high--
		} else {
			low++
		}
	}
}

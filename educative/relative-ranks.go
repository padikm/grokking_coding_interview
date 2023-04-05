package educative

import (
	"strconv"
)

//
func MaxHeap(a []int) []int {
	for i := len(a) / 2; i >= 0; i-- {
		maxifyHeap(a, i)
	}
	return a
}

func maxifyHeap(a []int, i int) {
	left := 2 * i
	right := (2 * i) + 1
	maxone := i
	if left < len(a) && a[left] > a[maxone] {
		maxone = left
	}

	if right < len(a) && a[right] > a[maxone] {
		maxone = right
	}

	if maxone != i {
		a[maxone], a[i] = a[i], a[maxone]
		maxifyHeap(a, maxone)
	}
}

func PopMaxHeap(a []int) (int, []int) {
	maxOne := a[0]
	a[0] = a[len(a)-1]
	maxifyHeap(a, 0)
	a = a[:len(a)-1]
	return maxOne, a
}

func FindRelativeRanks(a []int) []string {
	m := make(map[int]int)
	for i := 0; i < len(a); i++ {
		m[a[i]] = i
	}
	res := make([]string, len(a))
	a = MaxHeap(a)
	//
	i := 1
	item := 0
	for len(a) > 0 {
		item, a = PopMaxHeap(a)
		if i == 1 {
			res[m[item]] = "Gold Medal"
		} else if i == 2 {
			res[m[item]] = "Silver Medal"
		} else if i == 3 {
			res[m[item]] = "Bronze Medal"
		} else {
			res[m[item]] = strconv.Itoa(i)
		}
		i++
	}
	return res
}

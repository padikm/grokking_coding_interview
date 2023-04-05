package educative

import (
	"sort"
)

func TripletSumClosetoTarget(a []int, k int) int {
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		if searchCloser(a, a[i], i, -(k - 1)) {
			return k - 1
		}
		k--
	}
	return k
}

func searchCloser(a []int, start, i, k int) bool {
	left := i + 1
	right := len(a) - 1
	for left < right {
		sum := a[left] + a[right] + start + k
		if sum == 0 {
			left++
			right--
			return true
		} else if sum > 0 {
			right--
		} else {
			left++
		}
	}
	return false
}

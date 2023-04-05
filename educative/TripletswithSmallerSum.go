package educative

import "sort"

func TripletswithSmallerSum(a []int, k int) int {

	sort.Ints(a)
	count := 0
	for i := 0; i < len(a); i++ {
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		left := i + 1
		right := len(a) - 1
		j := a[i]
		for left < right {
			sum := k - a[left] - a[right] - j
			if sum > 0 {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}
	return count

}

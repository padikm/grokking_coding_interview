package leet

import "sort"

func Boatstosavepeople(a []int, k int) int {
	sort.Ints(a)
	l := 0
	r := len(a) - 1
	count := 0
	for l <= r {
		sum := a[l] + a[r]
		if sum <= k {
			count++
			l++
			r--
		} else {
			if a[r] <= sum {
				count++
			}
			r--
		}
	}
	return count
}

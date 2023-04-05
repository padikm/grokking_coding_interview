package educative

import (
	"sort"
)

func TripletSumtoZero(a []int) [][]int {
	res := [][]int{}
	m := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		m[a[i]] = false
	}
	sort.Ints(a)
	for i := 0; i <= len(a)-3; i++ {
		if m[a[i]] {
			continue
		}
		m[a[i]] = true
		res = append(res, searRemainSum(a, -a[i], i)...)

	}
	return res
}

func searRemainSum(a []int, sum int, k int) [][]int {

	left := k + 1
	res := [][]int{}
	right := len(a) - 1

	for left < right {
		if a[left]+a[right] == sum {
			res = append(res, []int{-sum, a[left], a[right]})
			left++
			right--
			for left < right && a[left] == a[left-1] {
				left++
			}
			for left < right && a[right] == a[right+1] {
				right--
			}
		} else if a[right]+a[left] < sum {
			left++
		} else {
			right--
		}
	}

	return res
}

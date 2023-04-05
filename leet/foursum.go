package leet

import "sort"

func FourSum(a []int, k int) [][]int {

	sort.Ints(a)
	var res [][]int
	for j := 0; j <= len(a)-4; j++ {
		if j > 0 && a[j] == a[j-1] {
			continue
		}
		//tArray := a[j+1 : len(a)]
		for i := j + 1; i <= len(a)-3; i++ {
			if i > j+1 && a[i] == a[i-1] {
				continue
			}
			res = append(res, twoSum(a, a[i], a[j], i+1, k)...)
		}
	}
	return res
}

func twoSum(a []int, f, s int, l int, k int) [][]int {
	var res [][]int
	r := len(a) - 1
	for l < r {
		sum := a[l] + a[r] + f + s
		if sum == k {
			res = append(res, []int{s, f, a[l], a[r]})
			l++
			r--
			for l < len(a)-1 && a[l] == a[l-1] {
				l++
			}
			for r > 0 && a[r] == a[r+1] {
				r--
			}
		} else if sum < k {
			l++
		} else {
			r--
		}
	}
	return res
}

package leet

import (
	"sort"
)

func MinimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	l := 0
	h := nums[len(nums)-1] - nums[0]
	count := 0
	for l < h {
		mid := (l + h) / 2
		count = 0
		for i := 0; i+1 < len(nums); i++ {
			if mid >= nums[i+1]-nums[i] {
				i++
				count++
			}
		}
		if count >= p {
			h = mid
		} else {
			l = mid + 1
		}
	}
	return minimizeMax(nums, p)
	//return l
}

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	l := 0
	h := nums[len(nums)-1] - nums[0]
	count := 0
	for h != l {
		//mid := (l + h) / 2
		count = 0
		for i := 0; i+1 < len(nums); i++ {
			if l >= nums[i+1]-nums[i] {
				i++
				count++
			}
		}
		if count >= p {
			return l
		}
		l++

	}
	return l
}

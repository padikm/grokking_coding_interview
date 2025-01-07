package slidingwindow

import "math"

func LongestSubarray(nums []int) int {
	return longestSubarray(nums)
}
func longestSubarray(nums []int) int {
	res := 0
	k := 1
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			k--
		}
		for k < 0 {
			if nums[j] == 0 {
				k++
			}
			j++
		}
		res = int(math.Max(float64(res), float64(i-j)))
	}
	return res
}

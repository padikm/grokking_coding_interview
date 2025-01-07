package slidingwindow

import "slices"

func maxFrequency(nums []int, k int) int {
	var sum, res, j int
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		for nums[i]*(i-j+1)-sum > k {
			sum = sum - nums[j]
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

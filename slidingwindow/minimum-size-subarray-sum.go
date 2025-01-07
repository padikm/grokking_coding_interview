package slidingwindow

import "math"

func MinSubArrayLen(k int, nums []int) int {
	j := 0
	res := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		k = k - nums[i]
		for k <= 0 {
			res = int(math.Min(float64(res), float64(i-j+1)))			
			k = k + nums[j]
			j++
		}
	}

	return res %(len(nums)+1)
}
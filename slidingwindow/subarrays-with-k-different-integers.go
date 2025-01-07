package slidingwindow

import "math"

func SubarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithKDistinct1(nums, k)
}

func subarraysWithKDistinct1(nums []int, k int) int {
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

	 if res % (len(nums) + 1) == 0 {
		return -1
	 }
	 return res % (len(nums)+1)
}

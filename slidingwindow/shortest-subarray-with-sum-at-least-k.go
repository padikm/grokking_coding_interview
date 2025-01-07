package slidingwindow

import "math"

func ShortestSubarray(nums []int, k int) int {
	return shortestSubarray1(nums, k)
}

//2,-1,2,1
func shortestSubarray1(nums []int, k int) int {
	res := len(nums) + 1
	sum := 0
	j := 0

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum >= k {
			
				res = int(math.Min(float64(res), float64(i-j+1)))
			for j <= i {
				sum = sum - nums[j]
				if sum >= k {
					res = int(math.Min(float64(res), float64(i-j)))
				}
				j++
			}
		}

	}
	return res % (len(nums) + 1)
}

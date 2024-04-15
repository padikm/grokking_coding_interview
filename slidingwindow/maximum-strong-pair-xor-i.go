package slidingwindow

import (
	"math"
	"sort"
)

func MaximumStrongPairXor(nums []int) int {
	return maximumStrongPairXor1(nums)
}

func maximumStrongPairXor1(nums []int) int {
	sort.Ints(nums)
	xor := 0
	for j := 0; j < len(nums); j++ {
		i := j
		for i < len(nums) && math.Abs(float64(nums[j])-float64(nums[i])) <= math.Min(float64(nums[j]), float64(nums[i])) {
			xor = int(math.Max(float64(xor), float64(nums[j]^nums[i])))
			i++
		}
	}
	return xor
}

package slidingwindow

import "math"

func LongestOnes(nums []int, k int) int {
	return longestOnes1(nums, k)
}

func longestOnes1(nums []int, k int) int {
	j := 0
	res := 0
	oneCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			oneCount++
		}
	}

	if oneCount+k >= len(nums) {
		return len(nums)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && k >= 0 {
			k--
		}
		if k == 0 {
			res = int(math.Max(float64(res), float64(i-j+1)))
		}
		for k < 0 {
			if nums[j] == 0 {
				k++
			}
			j++
		}
	}
	return res
}

func longestOnes2(nums []int, k int) int {
	var z, res, j int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			z++
		}

		for z > k {
			if nums[j] == 0 {
				z--
			}
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

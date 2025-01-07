package slidingwindow

import "fmt"

func maximumSubarraySum(nums []int, k int) int64 {
	j := 0
	m := make(map[int]int)
	var sum, res int64
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		sum = sum + int64(nums[i])
		for m[nums[i]] > 1 && j <= i {
			m[nums[j]]--
			sum = sum - int64(nums[j])
			j++
		}
		if i-j+1 == k {
			res = max64(res, sum)
			sum = sum - int64(nums[j])
			m[nums[j]]--
			fmt.Println(nums[j : i+1])
			j++
		}
	}
	return res
}

func max64(i, j int64) int64 {
	if i > j {
		return i
	}

	return j
}

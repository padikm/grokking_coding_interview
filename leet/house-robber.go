package leet

func Rob(nums []int) int {
	return robUtil(nums, 0, 0, []int{})
}

//recursion
func robUtil(nums []int, i int, res int, dp []int) int {
	if len(dp) == 0 {
		dp = make([]int, len(nums))
	}

	if i >= len(nums) {
		return res
	}
	if dp[i] != 0 {
		return dp[i]
	}
	sum := 0

	for j := i; j < len(nums); j++ {
		sum = max(robUtil(nums, j+2, res+nums[j], dp), sum)
	}
	dp[i] = res
	return sum
}

//DP?

func robUtilDp(nums []int, i int, dp []int) int {
	if len(dp) == 0 {
		dp = make([]int, len(nums))
		for j := 0; j < len(dp); j++ {
			dp[j] = -1
		}
	}

	if i >= len(nums) {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}

	dp[i] = max(robUtilDp(nums, i+2, dp)+nums[i], robUtilDp(nums, i+1, dp))
	return dp[i]
}

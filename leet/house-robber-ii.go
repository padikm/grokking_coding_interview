package leet

func RobCircle(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rodCutUtil(nums, 0, []int{}, len(nums)-1), rodCutUtil(nums, 1, []int{}, len(nums)))
}

func rodCutUtil(nums []int, i int, dp []int, n int) int {

	if len(dp) == 0 {
		dp = make([]int, len(nums))
		for i := 0; i < len(nums); i++ {
			dp[i] = -1
		}
	}
	if i >= n {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}

	dp[i] = max(rodCutUtil(nums, i+2, dp, n)+nums[i], rodCutUtil(nums, i+1, dp, n))
	return dp[i]
}

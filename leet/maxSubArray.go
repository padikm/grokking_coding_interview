package leet

//func MaxSubArray(nums []int) int {
//	maxVal := -2147483648
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j <= len(nums); j++ {
//			maxVal = max(maxVal, sum(nums[i:j]))
//		}
//	}
//	return maxVal
//}
//func sum(num []int) int {
//	var sumI int
//	for _, j := range num {
//		sumI = sumI + j
//	}
//	return sumI
//}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSubArray(a []int) int {
	minNow := a[0]
	minFinal := a[0]
	for i := 1; i < len(a); i++ {
		if minNow > 0 {
			minNow = a[i] + minNow
		} else {
			minNow = a[i] + 0
		}
		minFinal = max(minFinal, minNow)
	}
	return minFinal
}

func MaxSubArrayDp(nums []int) int {
	var dp [10]int
	dp[0] = nums[0]
	maxI := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i] + 0
		}
		maxI = max(dp[i], maxI)
	}
	return maxI
}

package leet

func UniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return UniquePathsUtil(m, n, [][]int{})
}

func UniquePathsUtil(m int, n int, dp [][]int) int {
	if len(dp) == 0 {
		dp = make([][]int, m)
		for i := 0; i < len(dp); i++ {
			dp[i] = make([]int, n)
		}
	}

	if m == 1 || n == 1 {
		return 1
	}

	if dp[m-1][n-1] != 0 {
		return dp[m-1][n-1]
	}
	dp[m-1][n-1] = UniquePathsUtil(m, n-1, dp) + UniquePathsUtil(m-1, n, dp)
	return dp[m-1][n-1]
}

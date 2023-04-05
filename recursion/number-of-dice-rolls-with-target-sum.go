package recursion

func NumRollsToTarget(n, k int, target int) int {
	return numRollsToTarget(n, k, target, 0)

}
func numRollsToTarget(n, k int, target int, count int) int {
	tempRes := 0
	if target == 0 && count == n {
		return 1
	}
	for i := 1; i <= k && i <= target; i++ {
		tempRes = tempRes + numRollsToTarget(n, k, target-i, count+1)
	}
	return tempRes
}

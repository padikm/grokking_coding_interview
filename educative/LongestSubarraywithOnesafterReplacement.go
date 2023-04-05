package educative

func LongestSubarraywithOnesafterReplacement(a []int, k int) int {
	var maxLen, maxCount, wStart, res int
	for i := 0; i < len(a); i++ {
		maxLen++
		if a[i] == 0 {
			maxCount++
		}
		for maxCount > k {
			if res < maxLen {
				res = maxLen - 1
			}
			if a[wStart] == 0 {
				maxCount--
			}
			wStart++
			maxLen--
		}
		if res < maxLen {
			res = maxLen
		}
	}

	return res

}

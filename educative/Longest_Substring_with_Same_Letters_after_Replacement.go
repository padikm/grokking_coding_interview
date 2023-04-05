package educative

func Longest_Substring_with_Same_Letters_after_Replacement(s string, k int) int {
	window_start := 0
	maxCount := 0
	res := 0
	maxLen := 0
	m := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		m[r]++
		maxLen++
		if m[r] > maxCount {
			maxCount = m[r]
		}
		for i-window_start-maxCount+1 > k {
			if res < maxLen {
				res = maxLen - 1
				maxLen = 0

			}
			rs := rune(s[window_start])
			m[rs]--
			window_start++
		}
		if maxLen < i-window_start+1 {
			maxLen = i - window_start + 1
		}
	}
	return maxLen

}

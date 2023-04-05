package educative

func LongestSubstringWithKDistChar(s string, k int) int {
	start := 0
	maxLen := 0
	res := 0
	var flag = true
	m := make(map[rune]int)
	for _, e := range s {
		m[e] = m[e] + 1
		maxLen++
		for len(m) > k {
			flag = false
			m[rune(s[start])] = m[rune(s[start])] - 1
			if m[rune(s[start])] == 0 {
				delete(m, rune(s[start]))
			}
			if maxLen > res {
				res = maxLen - 1
			}
			maxLen--
			start++
		}
	}
	if flag {
		return maxLen
	}
	return res

}

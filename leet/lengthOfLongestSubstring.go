package leet

func LengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start := 0
	count := 0
	maxLen := 0
	//flag := false
	for _, e := range s {
		m[e]++
		count++
		for m[e] != 1 {
			r := rune(s[start])
			if count > maxLen {
				maxLen = count - 1
			}
			count--
			m[r]--
			start++
		}
	}
	if count > maxLen {
		return count
	}
	return maxLen
}

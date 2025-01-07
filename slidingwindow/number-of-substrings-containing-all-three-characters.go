package slidingwindow

func NumberOfSubstrings(s string) int {
	return numberOfSubstrings1(s, 3) - numberOfSubstrings1(s, 2)
}

func numberOfSubstrings1(s string, k int) int {
	res := 0
	j := 0
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
		for len(m) > k {
			m[string(s[j])]--
			if m[string(s[j])] == 0 {
				delete(m, string(s[j]))
			}
			j++
		}
		res = res + i - j + 1

	}
	return res
}

package slidingwindow

func MaximumLengthSubstring(s string) int {

	j := 0
	res := 0	
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		s1 := string(s[i])
		m[s1]++		
		for m[s1] > 2 {
			m[string(s[j])]--
			j++
		}
		if m[s1] <= 2 {
			res = max(res, i-j+1)
		}

	}
	return res
}

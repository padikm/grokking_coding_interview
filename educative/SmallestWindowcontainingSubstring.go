package educative

func SmallestWindowcontainingSubstring(s string, p string) string {
	m := make(map[rune]int)
	for i := 0; i < len(p); i++ {
		m[rune(p[i])]++
	}
	count := 0
	res := 0
	start := 0
	var minLen string
	maxConut := len(s)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		res++
		if _, ok := m[r]; ok {
			m[r]--
			if m[r] == 0 {
				count++
			}
		}

		if count == len(m) && res < maxConut {
			maxConut = res
			minLen = s[start : i+1]
		}

		for count == len(m) && start < len(s) {
			r := rune(s[start])
			res--
			if v, ok := m[r]; ok && v == 0 {
				count--
			}
			if res < maxConut {
				maxConut = res
				minLen = s[start : i+1]
			}
			if _, ok := m[r]; ok {
				m[r]++
			}
			start++
		}

	}
	return minLen

}

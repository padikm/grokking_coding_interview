package educative

func StringAnagrams(s string, p string) []int {
	m := make(map[rune]int)
	for _, v := range p {
		m[v]++
	}
	count := 0
	start := 0
	res := make([]int, 0)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if _, ok := m[r]; ok {
			m[r]--
			if m[r] == 0 {
				count++
			}
		}
		if count == len(m) {
			res = append(res, start)
			//t := rune(s[start])
			//m[t]++
			//start++
			//count--
			//continue
		}
		if i >= len(p)-1 {
			v := rune(s[start])
			start++
			if _, ok := m[v]; ok {
				if m[v] == 0 {
					count--
				}
				m[v]++
			}
		}
	}
	return res
}

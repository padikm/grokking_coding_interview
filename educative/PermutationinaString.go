package educative

func PermutationinaString(s string, p string) bool {

	mp := make(map[rune]int)
	start := 0
	count := 0
	for i := 0; i < len(p); i++ {
		r := rune(p[i])
		mp[r]++
	}
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if _, ok := mp[r]; ok {
			if _, ok := mp[r]; ok {
				mp[r]--
				if mp[r] == 0 {
					count++
				}
			}

		}
		if count == len(mp) {
			return true
		}
		if i >= len(p)-1 {
			s := rune(s[start])
			start++
			if v, ok := mp[s]; ok {
				if v == 0 {
					count--
				}
				mp[s]++
			}
		}
	}
	return false
}

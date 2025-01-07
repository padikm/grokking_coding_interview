package slidingwindow

import "math"

func CharacterReplacement(s string, k int) int {
	res := 0
	j := 0
	maxOfCnt := 0
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
		maxOfCnt = max(maxOfCnt, m[string(s[i])])
		for i-j-maxOfCnt+1 > k {
			m[string(s[j])]--
			j++
		}

		if i-j-maxOfCnt+1 <= k {
			res = int(math.Max(float64(res), float64(i-j+1)))
		}
	}
	return res
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

package slidingwindow

import "math"

func BalancedString(s string) int {
	return balancedString1(s)
}

// WWQQRRRRQRQQ
func balancedString1(s string) int {
	res := 0
	b := len(s) / 4
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}
	j := 0
	res = len(s)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]--
		for j< len(s) && isBalanced(m, b) {
			res = int(math.Min(float64(res),float64(i-j + 1)))
			m[string(s[j])]++
			j++
		}
		
	}
	return res
}

func isBalanced(m map[string]int, b int) bool {
	for _, v := range m {
		if v > b {
			return false

		}
	}
	return true
}

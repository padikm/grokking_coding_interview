package slidingwindow

import (
	"strings"
)

func CountGoodSubstrings(s string) int {
	return countGoodSubstrings1(s)
}

func countGoodSubstrings1(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		j := i
		for j < len(s) && (j-i+1) <= 3 && i+3 <= len(s) {
			if strings.Count(s[i:i+3], s[j:j+1]) > 1 {			
				break
			}
			j++
		}
		if i + 3 == j {
			res++
		}
	}
	return res
}

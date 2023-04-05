package geekForGeek

import "fmt"

func CountPalindromicSubsequenceGivenString() {
	m := make(map[string]int)
	fmt.Println(countPalindromicSubsequenceGivenString("abcdabcdabcdabc", 0, "", 0, m))
	fmt.Println(len(m))
}

func countPalindromicSubsequenceGivenString(s string, j int, res string, count int, m map[string]int) int {
	if j == len(s) {
		return count
	}
	for i := j; i < len(s); i++ {
		if isPalindrome(res + string(s[i])) {
			m[res+string(s[i])]++
			//fmt.Println(res + string(s[i]))
			count = count + 1
		}
		count = countPalindromicSubsequenceGivenString(s, i+1, res+string(s[i]), count, m)
	}
	return count
}

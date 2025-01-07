package dp

func CountSubstrings(s string) int {
	m := make(map[string]bool)
	return countSubstrings1(s, m)
}

func countSubstrings1(s string, m map[string]bool) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if m[string(s[i:j+1])] {
				res++
				continue
			}
			if isPalindrome(string(s[i : j+1])) {
				m[string(s[i:j+1])] = true
				res++
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

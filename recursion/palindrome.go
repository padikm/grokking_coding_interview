package recursion

import (
	"strings"
)

func Palindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] == s[len(s)-1] {
		return Palindrome(s[1 : len(s)-1])
	}
	return false
}

func IsPalindrome(s string) bool {
	var res string
	lower := strings.ToLower(s)
	for i := range lower {
		if (lower[i] >= 97 && lower[i] <= 122) || (lower[i] >= 48 && lower[i] <= 57) {
			res = res + string(lower[i])
		}
	}
	return Palindrome(res)
}

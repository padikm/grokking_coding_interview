package geekForGeek

import "fmt"

func PrintAllPalPartitionOfStr(s string) {
	printAllPalPartitionOfStr(s, 0, "")
}

func printAllPalPartitionOfStr(s string, j int, res string) {
	if j == len(s) {
		return
	}
	for i := j; i < len(s); i++ {
		if isPalindrome(string(s[j : i+1])) {
			fmt.Println(string(s[j : i+1]))
			printAllPalPartitionOfStr(s, i+1, string(s[j:i+1]))
		}
	}
}

func isPalindrome(s string) bool {
	r := reverseString(s)
	if s == r {
		return true
	}
	return false
}

func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _, v := range str {
		result = string(v) + result
	}
	return
}

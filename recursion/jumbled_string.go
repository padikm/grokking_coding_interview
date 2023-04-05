package recursion

import "fmt"

func JumbledString(a []string, s string) {
	fmt.Println(jumbledString(a, "", 0, s, 0))
}

func jumbledString(a []string, res string, j int, s string, count int) int {
	if res == s {
		fmt.Println(res)
		return count + 1
	}
	if len(res) == len(s) {
		return count
	}
	for i := j; i < len(a); i++ {
		count = jumbledString(a, res+a[i], i+1, s, count)
	}
	return count
}

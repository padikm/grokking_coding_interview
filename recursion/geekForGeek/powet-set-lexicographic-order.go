package geekForGeek

import "fmt"

func PowerSetLexicographicOrder(s []string) {
	powerSetLexicographicOrder(s, len(s), 0, "")
}

func powerSetLexicographicOrder(s []string, k int, j int, res string) {
	if len(res) >= k {
		//fmt.Println(res)
		return
	}
	for i := j; i < len(s); i++ {
		fmt.Println(res + s[i])
		powerSetLexicographicOrder(s, k, i+1, res+s[i])
	}
}

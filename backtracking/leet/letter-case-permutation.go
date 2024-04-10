package leet

import (
	"fmt"
	"strings"
)

func LetterCasePermutation(s string) []string {
	return letterCasePermutation1(s, "", 0)
}

func letterCasePermutation1(s string, res string, j int) []string {
	if len(res) == len(s) {
		fmt.Println(res)
		return []string{}
	}
	for i := 0; i < len(s); i++ {
		res = s[0:i] + strings.ToUpper(string(s[i])) + s[i+1:]
		//res = res + string(s[i])
		letterCasePermutation1(s, res, i+1)
		res = res[0:i] + strings.ToLower(string(res[i])) + res[i+1:]
		//res = res[:len(res)-1]
	}
	return []string{}
}

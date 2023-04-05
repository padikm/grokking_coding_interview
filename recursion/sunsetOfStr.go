package recursion

import "fmt"

func SubSetOfStr(s string, r string) {
	if len(s) == 0 {
		fmt.Println(r)
		return
	}
	SubSetOfStr(s[1:len(s)], r+string(s[0]))
	SubSetOfStr(s[1:len(s)], r)
}

func SubSetOfStrReturn(s string, r string) []string {
	if len(s) == 0 {
		return []string{r}
	}
	res := make([]string, 0)
	res = append(res, SubSetOfStrReturn(s[1:len(s)], r+string(s[0]))...)
	res = append(res, SubSetOfStrReturn(s[1:len(s)], r)...)
	return res
}

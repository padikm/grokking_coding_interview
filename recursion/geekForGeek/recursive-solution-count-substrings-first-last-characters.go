package geekForGeek

import "fmt"

func CountSubStrSameFirstLastChar(s string) {
	m := make(map[string]bool)
	countSubStrSameFirstLastChar(s, "", m)
}

func countSubStrSameFirstLastChar(s string, res string, m map[string]bool) {
	if len(s) == 0 {
		if len(res) > 0 && res[0] == res[len(res)-1] {
			if !m[res] {
				fmt.Println(res)
				m[res] = true
			}
		}

		return
	}

	countSubStrSameFirstLastChar(s[1:len(s)], res+string(s[0]), m)
	countSubStrSameFirstLastChar(s[1:len(s)], res, m)
}

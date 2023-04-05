package leet

func Backspacestringcompare(s, t string) bool {
	if getAfterBSpace(s) == getAfterBSpace(t) {
		return true
	}
	return false
}

func getAfterBSpace(s string) string {
	var count int
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])
		if c == '#' {
			count++
		} else {
			if count > 0 {
				count--
			} else {
				res = string(c) + res
			}
		}
	}
	return res
}

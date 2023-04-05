package recursion

func SkipCharFromStr(s string, c string) string {
	return skipCharFromStr(s, c)
}

func skipCharFromStr(s string, c string) string {
	if len(s) == 0 {
		return s
	}
	if string(s[0]) != c {
		return string(s[0]) + skipCharFromStr(s[1:len(s)], c)
	}
	return skipCharFromStr(s[1:len(s)], c)
}

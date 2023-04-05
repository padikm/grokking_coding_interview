package recursion

func ReverseString(s string) string {

	if len(s) <= 1 {
		return s
	}
	return ReverseString(s[1:len(s)]) + string(s[0])
}

func ReverseStringInPlace(s []byte) {

	if len(s) <= 1 {
		return
	}
	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	ReverseStringInPlace(s[1 : len(s)-1])
}

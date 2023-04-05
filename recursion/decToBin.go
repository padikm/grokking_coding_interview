package recursion

import "strconv"

func ConvertDecToBin(n int, s *string) {
	if n <= 0 {
		return
	}
	*s = strconv.Itoa(n%2) + *s
	n = n / 2
	ConvertDecToBin(n, s)
}

func ConvertDecToBinReturn(n int) string {
	if n <= 0 {
		return ""
	}
	return ConvertDecToBinReturn(n/2) + strconv.Itoa(n%2)
}

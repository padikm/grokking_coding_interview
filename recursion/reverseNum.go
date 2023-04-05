package recursion

import "math"

func ReverseNumber(n int) int {
	return reverseNum(n, 0)
}

func reverseNum(n int, i int) int {
	if n <= 9 {
		return n
	}
	digit := int(math.Log10(float64(n)))
	return (n%10)*int(math.Pow(float64(10), float64(digit))) + reverseNum(n/10, i+1)
}

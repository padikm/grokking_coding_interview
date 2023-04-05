package recursion

func SumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + SumOfDigits(n/10)
}

func AddDigits(n int) int {
	if n >= 0 && n <= 9 {
		return n
	}
	return AddDigits(n%10 + AddDigits(n/10))
}

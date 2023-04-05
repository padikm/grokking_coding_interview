package recursion

func CountZeros(n int) int {
	if n == 0 {
		return 0
	}
	if n%10 == 0 {
		return 1 + CountZeros(n/10)
	}
	return CountZeros(n / 10)
}

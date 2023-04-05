package recursion

func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func FibOpt(n int, m map[int]int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if i, ok := m[n]; ok {
		return i
	}
	res := FibOpt(n-1, m) + FibOpt(n-2, m)
	m[n] = res
	return res
}

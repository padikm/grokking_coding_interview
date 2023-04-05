package educative

func DutchNationalFlagProblem(a []int) []int {
	l := 0
	r := len(a) - 1

	for i := 0; i < len(a); i++ {
		for l < len(a) && a[l] == 0 {
			l++
		}
		for r >= 0 && a[r] == 2 {
			r--
		}
		if i < r && a[i] == 2 {
			swap(a, i, r)
			r--
		} else if i > l && a[i] == 0 {
			swap(a, i, l)
			l++
		}
		for l <= i && a[i] < a[l] {
			swap(a, l, i)
			l++
		}
	}
	return a
}

func swap(a []int, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

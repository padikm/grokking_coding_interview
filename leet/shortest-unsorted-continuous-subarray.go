package leet

func Shortestunsortedcontinuoussubarray(a []int) int {
	l := 0
	r := len(a) - 1
	for l <= r {
		if l > 0 && a[l] < a[l-1] {
			l = l - 1
			break
		}
		l++
	}

	for r > 0 {
		if a[r-1] > a[r] {
			break
		}
		r--
	}
	if r < 0 {
		r = len(a) - 1
	}
	if l > r {
		return 0
	}
	low := findLow(a[l : r+1])
	high := findHigh(a[l : r+1])
	lowPos := findNextLowPos(a, l-1, low)
	highPos := findNextHighPos(a, r+1, high)
	//fmt.Println(lowPos, " ", highPos)
	//fmt.Println(a[lowPos+1 : highPos+1])
	return len(a[lowPos+1 : highPos+1])
}

func findLow(a []int) int {
	l := a[0]
	for _, v := range a {
		if l > v {
			l = v
		}
	}

	return l
}

func findNextLowPos(a []int, pos, k int) int {
	for i := pos; i >= 0; i-- {
		if k >= a[i] {
			return i
		}
	}
	return -1
}

func findNextHighPos(a []int, pos, k int) int {
	for i := pos; i < len(a); i++ {
		if k <= a[i] {
			return i - 1
		}
	}
	return len(a) - 1
}

func findHigh(a []int) int {
	h := a[0]
	for _, v := range a {
		if h < v {
			h = v
		}
	}
	return h
}

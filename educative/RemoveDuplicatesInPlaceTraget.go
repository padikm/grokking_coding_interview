package educative

func RemoveDuplicatesInPlaceTarget(a []int, t int) []int {
	cur := 1
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != t {
			a[cur-1] = a[i]
			cur++
		} else {
			count++
		}
	}
	return a[0 : len(a)-count]
}

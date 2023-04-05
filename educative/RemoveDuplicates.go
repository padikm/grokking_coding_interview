package educative

func RemoveDuplicates(a []int) int {
	cur := a[0]
	count := 1
	for i := 1; i < len(a); i++ {
		if cur != a[i] {
			count++
		}
		cur = a[i]
	}
	return count

}

func RemoveDuplicatesInPlace(a []int) []int {
	cur := 0
	for i := 1; i < len(a); i++ {
		if a[cur] != a[i] {
			a[cur+1] = a[i]
			cur++
		}
	}
	return a[0 : cur+1]

}

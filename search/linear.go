package search

func LinearSearch(a []int, i int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

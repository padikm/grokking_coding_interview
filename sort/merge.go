package sort

func MergeSort(a []int) []int {
	if 2 > len(a) {
		return a
	}
	m := len(a) / 2
	f := MergeSort(a[:m])
	l := MergeSort(a[m:])
	return merge(f, l)
}

func merge(a, b []int) []int {
	var i, j int
	res := make([]int, 0)
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	for i < len(a) {
		res = append(res, a[i])
		i++
	}
	for j < len(b) {
		res = append(res, b[j])
		j++
	}
	return res
}

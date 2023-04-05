package recursion

func LinearSearch(a []int, k int) bool {
	return linearSearch(a, 0, k)
}

func LinearSearchAllIndex(a []int, k int) []int {
	return linearSearchAllIndexNoArgInOrder(a, 0, k)
}

func linearSearch(a []int, i, k int) bool {
	if i == len(a) {
		return false
	}
	return a[i] == k || linearSearch(a, i+1, k)
}

func linearSearchAllIndex(a []int, i, k int, res []int) []int {
	if i == len(a) {
		return res
	}
	if a[i] == k {
		res = append(res, i)
	}
	return linearSearchAllIndex(a, i+1, k, res)
}

func linearSearchAllIndexNoArg(a []int, i, k int) []int {
	if i == len(a) {
		return []int{}
	}
	if a[i] == k {
		return append(linearSearchAllIndexNoArg(a, i+1, k), i)
	}
	return linearSearchAllIndexNoArg(a, i+1, k)
}

func linearSearchAllIndexNoArgInOrder(a []int, i, k int) []int {
	if i == len(a) {
		return []int{}
	}
	res := make([]int, 0)
	if a[i] == k {
		res = append(res, i)
	}
	r := linearSearchAllIndexNoArgInOrder(a, i+1, k)
	res = append(res, r...)
	return res
}

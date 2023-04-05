package leet

func KLenLrgSum(a []int, k int) int {
	if k > len(a) {
		return -1
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum = sum + a[i]
	}
	j := 0
	tSum := sum
	for i := k; i < len(a); i++ {
		tSum = tSum - a[j] + a[i]
		if tSum > sum {
			sum = tSum
		}
		j++
	}
	return sum
}

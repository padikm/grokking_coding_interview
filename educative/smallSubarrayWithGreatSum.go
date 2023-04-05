package educative

func FindMinSubArray(a []int, k int) int {

	for i := 1; i <= len(a); i++ {
		if k <= slideSum(a, i) {
			return i
		}

	}
	return -1
}

func slideSum(a []int, k int) int {
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

func FindMinSubArrayOptimized(a []int, k int) int {
	sum := 0
	j := 0
	minLen := len(a)
	flag := true
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
		for sum >= k {
			sum = sum - a[j]
			if minLen > (i - j) {
				flag = false
				minLen = i - j + 1
			}
			j++
		}
	}
	if flag {
		minLen = 0
	}
	return minLen
}

package educative

func SubarrayProductLessThanK(a []int, k int) int {
	var count int
	for i := 0; i < len(a); i++ {
		p := a[i]
		if p < k {
			count++
		} else {
			continue
		}
		for j := i + 1; j < len(a); j++ {
			p = p * a[j]
			if p < k {
				//res = append(res, a[i:j+1])
				count++
			} else {
				break
			}
		}
	}
	return count
}

func SubarrayProductLessThanKOpt(a []int, k int) int {
	l := 0
	r := 0
	var count int
	p := 1
	for l < len(a) && r < len(a) {
		p = p * a[r]
		for l <= r && p >= k {
			p = p / a[l]
			l++
		}
		count = count + r - l + 1
		r++
	}
	return count
}

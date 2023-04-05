package leet

func Intersect(a []int, b []int) []int {
	var res = make([]int, 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				res = append(res, a[i])
				b = append(b[0:j], b[j+1:len(b)]...)
				break
			}
		}
	}
	return res
}

func Intersect1(a, b []int) []int {
	m := make(map[int]int)
	var small, big []int
	if len(a) >= len(b) {
		small = b
		big = a
	} else {
		small = a
		big = b
	}

	for i := 0; i < len(small); i++ {
		m[small[i]]++
	}
	res := make([]int, 0)
	for i := 0; i < len(big); i++ {
		if v, ok := m[big[i]]; ok && v != 0 {
			res = append(res, big[i])
			m[big[i]]--
		}
	}
	return res
}

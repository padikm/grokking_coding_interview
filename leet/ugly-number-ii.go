package leet

import "math"

func NthUglyNumberBrute(n int) int {
	count := 0
	res := 0
	m := make(map[int]bool)
	for i := 1; count != n; i++ {
		j := i
		for j != 1 {
			if e := m[j]; e {
				j = 1
				break
			}
			if j%5 == 0 {
				j = j / 5
			} else if j%3 == 0 {
				j = j / 3
			} else if j%2 == 0 {
				j = j / 2
			} else {
				break
			}

		}
		if j == 1 {
			m[i] = true
			res = i
			count++
		}
	}
	return res
}

func NthUglyNumber(n int) int {

	res := make([]int, 0)
	res = append(res, 1)
	two := 2
	itwo := 1
	three := 3
	iThree := 1
	five := 5
	iFive := 1
	for i := 1; i < n; i++ {
		min := int(math.Min(float64(two), math.Min(float64(three), float64(five))))
		res = append(res, min)
		if min == two {
			two = 2 * res[itwo]
			itwo++
		}
		if min == three {
			three = 3 * res[iThree]
			iThree++
		}
		if min == five {
			five = 5 * res[iFive]
			iFive++
		}
	}
	return res[len(res)-1]
}

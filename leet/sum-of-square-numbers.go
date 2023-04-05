package leet

import "math"

func Sumofsquarenumbers(k int) bool {
	l := 0
	r := int(math.Sqrt(float64(k)))

	for l <= r {
		sum := l*l + r*r
		if sum == k {
			return true
		} else if sum < k {
			l++
		} else {
			r--
		}
	}
	return false
}

package leet

import (
	"math"
	"sort"
)

func Minimizemaximumpairsuminarray(a []int) int {
	sort.Ints(a)
	l := 0
	r := len(a) - 1
	sum := 0
	for l < r {
		temp := a[l] + a[r]
		sum = int(math.Max(float64(sum), float64(temp)))
		l++
		r--
	}
	return sum
}

package leet

import (
	"fmt"
	"sort"
)

func Numberofsubsequencesthatsatisfythegivensumcondition(a []int, k int) int {
	sort.Ints(a)
	mod := 1000000007
	l := 0
	r := len(a) - 1
	res := 0
	var pow = make([]int, len(a))
	pow[0] = 1
	for i := 1; i < len(a); i++ {
		pow[i] = (pow[i-1] * 2) % mod
	}
	for l <= r {
		if a[l]+a[r] <= k {
			//diff := r - l
			res = (res + pow[r-l]) % mod
			l++
		} else {
			r--
		}
	}
	fmt.Println(res)
	return int(res)
}

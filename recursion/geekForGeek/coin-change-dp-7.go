package geekForGeek

import (
	"fmt"
)

func CoinChangeCount(c []int, k int) {
	fmt.Println(coinChangeCount1(c, k, 0, []int{}, 0, -1))
}

func coinChangeCount1(c []int, k int, res int, a []int, j int, min int) int {
	if res == k {
		fmt.Println(res, a, len(a))
		if min == -1 || min > len(a) {
			min = len(a)
		}
		return min
	}
	if res > k {
		return min
	}
	for i := j; i < len(c); i++ {
		min = coinChangeCount1(c, k, res+c[i], append(a, c[i]), i, min)

	}
	return min
}

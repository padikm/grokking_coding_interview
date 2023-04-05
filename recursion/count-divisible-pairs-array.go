package recursion

import (
	"fmt"
)

func CountDivisiblePairsArray(n, m int) {
	res := make([]int, 0)
	countDivisiblePairsArray(n, m, res)
}

func countDivisiblePairsArray(n, m int, res []int) {
	if len(res) == n {
		fmt.Println(res)
		return
	}
	for i := 1; i <= m; i++ {
		//if i
		//s := strconv.Itoa(i)
		if len(res) > 0 {
			e := res[len(res)-1]
			if e%i == 0 || i%e == 0 {
				res = append(res, i)
			}
		} else {
			res = append(res, i)
		}
		countDivisiblePairsArray(n, m, res)
		res = res[:len(res)-1]
	}
}

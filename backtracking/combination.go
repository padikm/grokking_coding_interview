package backtracking

import "fmt"

func Combination(n int, k int) {
	combination(n, k, []int{}, 0)
}

func combination(n, k int, res []int, j int) {
	if len(res) == k {
		fmt.Println(res)
		return
	}
	for i := 1; i <= n; i++ {
		res = append(res, i)
		combination(n, k, res, i+1)
		res = res[:len(res)-1]
	}
}

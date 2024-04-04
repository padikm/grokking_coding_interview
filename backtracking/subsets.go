package backtracking

import "fmt"

func Subsets(res, a []int, j int) {
	fmt.Println(res)
	for i := j; i < len(a); i++ {
		res = append(res, a[i])
		Subsets(res, a, i+1)
		res = res[:len(res)-1]
	}
}

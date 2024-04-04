package backtracking

import "fmt"

func Permutation(res, a []int, n int) {
	if n == len(res) {
		fmt.Println(res)
		return
	}

	for i := 0; i < len(a); i++ {
		b := 0
		for _, v := range res {
			if a[i] == v {
				b = 1
				break
			}
		}
		if b != 1 {
			res = append(res, a[i])
		} else {
			continue
		}
		Permutation(res, a, n)
		res = res[:len(res)-1]
	}
}

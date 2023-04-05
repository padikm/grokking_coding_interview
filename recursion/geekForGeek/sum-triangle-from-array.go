package geekForGeek

import "fmt"

func SumTriangleFromArray(a []int) {
	sumTriangleFromArray(a, 0)
}

func sumTriangleFromArray(a []int, i int) {
	if len(a) == 1 {
		fmt.Println(a)
		return
	}
	if i+1 == len(a) {
		fmt.Println(a[0 : len(a)-1])
		sumTriangleFromArray(a[0:len(a)-1], 0)
		return
	}

	a[i] = a[i] + a[i+1]
	sumTriangleFromArray(a, i+1)
}

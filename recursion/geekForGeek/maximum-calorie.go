package geekForGeek

import "fmt"

func MaximumCalorie(a []int, n int) {
	fmt.Println(maximumCalorie(a, n, n-1))
}

func maximumCalorie(a []int, n int, i int) int {
	if i == 0 {
		return a[0]
	}
	if i == 1 {
		return a[0] + a[1]
	}
	if i == 2 {
		return max(a[0]+a[2], max(a[0]+a[1], a[1]+a[2]))
	}

	return max(max(maximumCalorie(a, n, i-1), maximumCalorie(a, n, i-2)+a[i]), maximumCalorie(a, n, i-3)+a[i-1]+a[i])
}

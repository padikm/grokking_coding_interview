package geekForGeek

import "fmt"

func MaximumProfitSaleWines(a []int, n int) {
	fmt.Println(maximumProfitSaleWines(a, 0, len(a)-1, 1))
}

func maximumProfitSaleWines(a []int, i, j int, y int) int {
	if i >= j {
		return a[i] * y
	}

	return max((a[i]*y)+maximumProfitSaleWines(a, i+1, j, y+1), (a[j]*y)+maximumProfitSaleWines(a, i, j-1, y+1))
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

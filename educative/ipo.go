package educative

func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	a := make([][]int, 0)
	for i := 0; i < len(capital); i++ {
		a = append(a, []int{capital[i], profits[i]})
	}
	maxProfitHeap := make([][]int, 0)
	twoDimHeap(a)
	small := make([]int, 0)
	small, lrgProfit := make([]int, 0), make([]int, 0)
	for ; k != 0; k-- {
		for len(a) > 0 && a[0][0] <= w {
			small, a = popMinHeap(a)
			maxProfitHeap = pushMaxHeap(maxProfitHeap, small[0], small[1])
		}
		if len(maxProfitHeap) == 0 {
			break
		}
		lrgProfit, maxProfitHeap = popMaxHeap(maxProfitHeap)
		w = w + lrgProfit[1]

	}
	return w
}

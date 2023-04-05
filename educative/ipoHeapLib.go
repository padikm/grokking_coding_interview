package educative

import "container/heap"

func FindMaximizedCapitalIpo(k int, w int, profits []int, capital []int) int {

	var ipo = Ipo{}
	for i := 0; i < len(capital); i++ {
		ipo = append(ipo, []int{capital[i], profits[i]})
	}
	maxProfit := maxProfit{}
	heap.Init(&maxProfit)
	heap.Init(&ipo)
	for ; k != 0; k-- {
		for len(ipo) > 0 && ipo[0][0] <= w {
			small := heap.Pop(&ipo).([]int)
			heap.Push(&maxProfit, small)
		}
		if len(maxProfit) == 0 {
			break
		}
		w = w + heap.Pop(&maxProfit).([]int)[1]
	}
	return w
}

type Ipo [][]int

func (ipo *Ipo) Len() int {
	//TODO implement me
	return len(*ipo)
}

func (ipo *Ipo) Less(i, j int) bool {
	//TODO implement me
	return (*ipo)[i][0] < (*ipo)[j][0]
}

func (ipo *Ipo) Swap(i, j int) {
	(*ipo)[i], (*ipo)[j] = (*ipo)[j], (*ipo)[i]
}

func (ipo *Ipo) Push(x interface{}) {
	*ipo = append(*ipo, x.([]int))
}

func (ipo *Ipo) Pop() interface{} {
	small := (*ipo)[len(*ipo)-1]
	*ipo = (*ipo)[:len(*ipo)-1]
	return small
}

type maxProfit [][]int

func (m *maxProfit) Len() int {
	return len(*m)
}

func (m *maxProfit) Less(i, j int) bool {
	return (*m)[i][1] > (*m)[j][1]
}

func (m *maxProfit) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxProfit) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *maxProfit) Pop() interface{} {
	large := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return large
}

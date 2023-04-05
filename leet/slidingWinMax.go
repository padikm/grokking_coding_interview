package leet

import "container/heap"

func MaxSlidingWindow1(a []int, k int) []int {
	fp := 0
	sp := 0
	res := make([]int, 0)
	if len(a) > 1 {
		sp = 1
	} else {
		return a
	}
	for fp < len(a) || sp < len(a) {
		if fp < len(a) {
			max := a[fp]
			temp := fp
			if fp+k > len(a) {
				break
			}
			for fp < k+temp && fp < len(a) {
				if a[fp] > max {
					max = a[fp]
				}
				fp++
			}
			fp = temp + k - 1
			res = append(res, max)
		}
		if sp < len(a) {
			max := a[sp]
			temp := sp
			if sp+k > len(a) {
				break
			}
			for sp <= k+temp-1 && sp < len(a) {
				if a[sp] > max {
					max = a[sp]
				}
				sp++
			}
			sp = temp + k - 1
			res = append(res, max)
		}

	}
	return res

}

func MaxSlidingWindow(a []int, k int) []int {
	fp := 0
	var ah arrHeap
	heap.Init(&ah)
	res := make([]int, 0)
	for i := 0; i < len(a); i++ {
		heap.Push(&ah, []int{a[i], i})
		for i-fp+1 >= k {
			for len(ah) > 0 && ah[0][1] < fp {
				heap.Pop(&ah)
			}
			res = append(res, ah[0][0])
			if len(ah) > 0 && ah[0][0] == a[fp] && ah[0][1] == fp {
				heap.Pop(&ah)
			}
			//for len(ah) > 0 && ah[0][0] == a[fp] && ah[0][1] <= fp {
			//	heap.Pop(&ah)
			//}
			fp++
		}

	}

	return res
}

type arrHeap [][]int

func (a *arrHeap) Len() int {
	//TODO implement me
	return len(*a)
}

func (a *arrHeap) Less(i, j int) bool {
	//TODO implement me
	return (*a)[i][0] > (*a)[j][0]
}

func (a *arrHeap) Swap(i, j int) {
	//TODO implement me
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func (a *arrHeap) Push(x interface{}) {
	//TODO implement me
	*a = append(*a, x.([]int))
}

func (a *arrHeap) Pop() interface{} {
	//TODO implement me
	small := (*a)[0]
	*a = (*a)[:len(*a)-1]
	return small
}

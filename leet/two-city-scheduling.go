package leet

import (
	"fmt"
	"sort"
)

func TwoCitySchedCost(costs [][]int) int {
	diff := make([][]int, 0)
	for i := 0; i < len(costs); i++ {
		d := costs[i][0] - costs[i][1]
		diff = append(diff, []int{d, costs[i][0], costs[i][1]})
	}

	sort.Slice(diff, func(i, j int) bool {
		if diff[i][0] < diff[j][0] {
			return true
		}
		return false
	})

	fmt.Println(diff)
	res := 0
	for i := 0; i < len(diff); i++ {
		if i < len(diff)/2 {
			res = res + diff[i][1]
		} else {
			res = res + diff[i][2]
		}
	}
	return res
}

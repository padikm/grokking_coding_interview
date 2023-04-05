package educative

import (
	"fmt"
	"sort"
)

/*
	fmt.Println(educative.IntervalIntersection(
		[][]int{{1, 4}, {5, 6}, {7, 9}},
		[][]int{{3, 5}, {6, 7}, {8, 9}}))

	fmt.Println(educative.IntervalIntersection(
		[][]int{{2, 6}, {7, 9}, {10, 13}, {14, 19}, {20, 24}},
		[][]int{{1, 4}, {6, 8}, {15, 18}}))

	fmt.Println(educative.IntervalIntersection(
		[][]int{{1, 3}, {5, 9}},
		[][]int{}))

	fmt.Println(educative.IntervalIntersection(
		[][]int{{3, 5}, {9, 20}},
		[][]int{{4, 5}, {7, 10}, {11, 12}, {14, 15}, {16, 20}}))
*/

func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	firstList = append(firstList, secondList...)
	sort.SliceStable(firstList, func(i, j int) bool {
		if firstList[i][0] < firstList[j][0] {
			return true
		}
		return false
	})
	var res = make([][]int, 0)
	fmt.Println(firstList)
	j := 0
	for i := 1; i < len(firstList); i++ {
		if firstList[j][1] > firstList[i][0] && firstList[j][1] > firstList[i][1] {
			res = append(res, []int{firstList[i][0], firstList[i][1]})
		} else if firstList[j][1] >= firstList[i][0] {
			res = append(res, []int{firstList[i][0], firstList[j][1]})
			j = i
		} else {
			j = i
		}

	}
	return res
}

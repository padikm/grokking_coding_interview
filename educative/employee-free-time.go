package educative

import "sort"

/*
	fmt.Println(educative.EmployeeFreeTime([][][]int{{{1, 2}, {5, 6}}, {{1, 3}}, {{4, 10}}}))
	fmt.Println(educative.EmployeeFreeTime([][][]int{{{1, 3}, {6, 7}}, {{2, 4}}, {{2, 5}, {9, 12}}}))
	fmt.Println(educative.EmployeeFreeTime([][][]int{{{2, 3}, {7, 9}}, {{1, 4}, {6, 7}}}))
	fmt.Println(educative.EmployeeFreeTime([][][]int{{{3, 5}, {8, 10}}, {{4, 6}, {9, 12}}, {{5, 6}, {8, 10}}}))
	fmt.Println(educative.EmployeeFreeTime([][][]int{{{1, 3}, {6, 9}, {10, 11}}, {{3, 4}, {7, 12}}, {{1, 3}, {7, 10}}, {{1, 4}}, {{7, 10}, {11, 12}}}))
	fmt.Println(educative.EmployeeFreeTime([][][]int{{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, {{2, 3}, {4, 5}, {6, 8}}}))
	fmt.Println(educative.EmployeeFreeTime([][][]int{{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}}, {{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}}, {{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}}, {{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}}}))

*/

func EmployeeFreeTime(a [][][]int) [][]int {
	var res, as = make([][]int, 0), make([][]int, 0)
	for i := 0; i < len(a); i++ {
		as = append(as, a[i]...)
	}
	j := 0

	sort.SliceStable(as, func(i, j int) bool {
		if as[i][0] < as[j][0] {
			return true
		}
		return false
	})

	for i := 1; i < len(as); i++ {
		if as[j][1] > as[i][0] && as[j][1] > as[i][1] {
			continue
		}
		if as[j][1] < as[i][0] {
			res = append(res, []int{as[j][1], as[i][0]})
		}
		j = i
	}
	return res
}

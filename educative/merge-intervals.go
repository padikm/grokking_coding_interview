package educative

import "sort"

func MergeIntervals(intervals [][]int) [][]int {

	var res = make([][]int, 0)
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	res = append(res, intervals[0])

	j := 0
	//start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if end >= intervals[i][0] && end <= intervals[i][1] {
			end = intervals[i][1]
			res[j][1] = end
		} else if end < intervals[i][0] {
			res = append(res, intervals[i])
			end = intervals[i][1]
			j++
		}
	}
	return res
}

package educative

import "sort"

func Insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false

	})
	var res [][]int
	res = append(res, intervals[0])
	end := res[0][1]
	j := 0
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

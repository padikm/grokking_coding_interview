package leet

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	subsetsWithDupVoid(nums, []int{}, &res)
	return res

}

func subsetsWithDupHelper(nums, res []int) [][]int {
	fres := make([][]int, 0)
	if len(nums) == 0 {
		return [][]int{res}
	}
	fres = append(fres, subsetsWithDupHelper(nums[1:], append(res, nums[0]))...)
	for len(nums) > 1 && nums[1] == nums[0] {
		nums = nums[1:]
		//subsetsWithDupHelper(nums[1:], res)
		//return fres
	}
	fres = append(fres, subsetsWithDupHelper(nums[1:], res)...)

	return fres
}

func subsetsWithDupVoid(nums, res []int, finres *[][]int) {
	if len(nums) == 0 {
		tRes := make([]int, 0)
		for i := 0; i < len(res); i++ {
			tRes = append(tRes, res[i])
		}
		*finres = append(*finres, tRes)
		return
	}
	subsetsWithDupVoid(nums[1:], append(res, nums[0]), finres)
	for len(nums) > 1 && nums[1] == nums[0] {
		nums = nums[1:]
		//subsetsWithDupHelper(nums[1:], res)
		//return fres
	}
	subsetsWithDupVoid(nums[1:], res, finres)

}

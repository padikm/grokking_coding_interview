package leet

import (
	"sort"
)

func CanPartitionKSubsets(nums []int, k int) bool {
	m := make(map[int][][]int)
	sort.Ints(nums)
	canPartitionKSubsets1(nums, k, []int{}, m, 0)
	for _, v := range m {
		if len(v) == k {
			return true
		}
	}
	return false
}

func canPartitionKSubsets1(nums []int, k int, res []int, m map[int][][]int, sum int) {
	if len(nums) == 0 {
		m[sum] = append(m[sum], res)
		return
	}
	canPartitionKSubsets1(nums[1:], k, append(res, nums[0]), m, sum+nums[0])
	//for len(nums) > 1 && nums[1] == nums[0] {
	//	nums = nums[1:]
	//}
	canPartitionKSubsets1(nums[1:], k, res, m, sum)
}

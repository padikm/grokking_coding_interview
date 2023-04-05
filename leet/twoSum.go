package leet

//check this again
func TwoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		rem := target - nums[i]
		var index = make([]int, 0)
		index = append(index, i)
		tempRem := rem
		for j := i + 1; j < len(nums); j++ {
			if tempRem == nums[j] {
				return []int{i, j}
			}
			if (rem - nums[j]) < 0 {
				continue
			}
			index = append(index, j)
			rem = rem - nums[j]
			if rem == 0 {
				if len(index) > 2 {
					break
				}
				return index
			}
		}
	}
	return nil
}

func TwoSumMax(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		rem := target - nums[i]
		if v, ok := m[rem]; ok && i != v {
			return []int{i, v}
		}
	}
	return []int{}
}

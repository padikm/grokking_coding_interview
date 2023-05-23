package leet

func Jump(nums []int) int {

	j := len(nums) - 1
	var res int
	for j > 0 {
		isOk := false
		tempJ := j
		for i := j - 1; i >= 0; i-- {
			if nums[i] >= tempJ-i {
				j = i
				isOk = true
			}
		}
		if !isOk {
			return 0
		}
		res++
	}
	return res
}

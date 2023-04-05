package leet

func CanJump(nums []int) bool {
	return canJump(nums)
}

func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}
	k := len(nums) - 1
	res := false
	i := k
	for i > 0 {
		res = false
		for j := k - 1; j >= 0; j-- {
			if nums[j] >= k-j {
				res = true
				i, k = j, j
				break
			}
		}
		if !res {
			return res
		}
	}
	return res
}

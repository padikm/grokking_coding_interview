package leet

func SubsetXORSum(nums []int) int {
	final := 0
	subsetXORSum(nums, []int{}, 0, &final)
	return final
}

func subsetXORSum(nums []int, res []int, j int, final *int) {
	tmp := 0
	for i := 0; i < len(res); i++ {
		tmp = tmp ^ res[i]
	}
	*final = *final + tmp
	for i := j; i < len(nums); i++ {
		res = append(res, nums[i])
		subsetXORSum(nums, res, i+1, final)
		res = res[:len(res)-1]
	}
}

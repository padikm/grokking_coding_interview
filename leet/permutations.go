package leet

func Permute(nums []int) [][]int {
	res := make([][]int, 0)
	permute(nums, &res, 0)
	return res
}

func permute(nums []int, res *[][]int, j int) {
	if len(nums)-1 == j {
		n := make([]int, 0)
		for i := 0; i < len(nums); i++ {
			n = append(n, nums[i])
		}
		*res = append(*res, n)
		return
	}
	for i := j; i < len(nums); i++ {
		nums = swapNums(nums, i, j)
		//*res = append(*res, nums)
		permute(nums, res, j+1)
	}
}

func swapNums(nums []int, i, j int) []int {
	nums[j], nums[i] = nums[i], nums[j]
	n := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		n = append(n, nums[i])
	}
	return n
}

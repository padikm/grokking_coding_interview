package slidingwindow

func numSubarrayProductLessThanK(nums []int, k int) int {
	var j, p, res int
	p = 1
	for i := 0; i < len(nums); i++ {
		p = p * nums[i]
		for p >= k && j <= i {
			p = p / nums[j]
			j++
		}
		res = res + i - j + 1
	}
	return res
}

package slidingwindow

func CountCompleteSubarrays(nums []int) int {

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	return countCompleteSubarrays1(nums, len(m)) - countCompleteSubarrays1(nums, len(m)-1)

}

func countCompleteSubarrays1(nums []int, n int) int {
	res, j := 0, 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		for len(m) > n {
			m[nums[j]]--
			if m[nums[j]] == 0 {
				delete(m, nums[j])
			}
			j++
		}
		res = res + i - j + 1
	}
	return res
}

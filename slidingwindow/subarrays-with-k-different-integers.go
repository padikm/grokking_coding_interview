package slidingwindow

func SubarraysWithKDistinct(nums []int, k int) int {
    return subarraysWithKDistinct1(nums,k)
}

func subarraysWithKDistinct1(nums []int, k int) int {
    res := 0
	for i:=0;i<len(nums);i++ {
		j:=i
		m := make(map[int]bool)
		for j<len(nums){
			m[nums[j]] = true
			if len(m) == k {
				res++
			}

			if len(m) > k {
				break
			}
			j++
		}
		if j == len(nums) && res == 0{
			return res
		}
	}
	return res
}
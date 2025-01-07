package slidingwindow

func NumberOfSubarrays(nums []int, k int) int {
	return numberOfSubarrays2(nums, k) - numberOfSubarrays2(nums, k-1)
}

func numberOfSubarrays1(nums []int, k int) int {
	res := 0
	j := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			k--
			count = 0
		}
		for k == 0 {
			k = k + nums[j]%2
			j++
			count++
		}
		res = res + count
	}
	return res
}

func numberOfSubarrays2(nums []int, k int) int {
	var c, j, res int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			c++
		}

		for c > k {
			c = c - nums[j]%2
			j++
		}
		res = res + i - j + 1
	}
	return res
}

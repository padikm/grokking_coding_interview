package slidingwindow

func NumberOfSubarrays(nums []int, k int) int {
    return numberOfSubarrays1(nums,k)
}

func numberOfSubarrays1(nums []int, k int) int {
    res :=0
	for i:=0;i<len(nums);i++ {
		j:=i
		odd :=0
		for j<len(nums) {
			if nums[j] %2 !=0 {
				odd++
			}
			if odd == k {
				res++
			}

			if odd > k {
				break
			}
			j++
		}
	}
	return res
}
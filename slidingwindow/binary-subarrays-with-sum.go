package slidingwindow

func NumSubarraysWithSum(nums []int, goal int) int {
    return numSubarraysWithSum1(nums,goal)
}

func numSubarraysWithSum1(nums []int, goal int) int {
    res :=0

	for i:=0;i<len(nums);i++ {
		j := i
		sum := 0
		for j<len(nums) && sum <= goal{
			sum += nums[j]
			if sum == goal {
				res++
			}
			j++
		}
	}
	return res
}
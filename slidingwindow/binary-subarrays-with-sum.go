package slidingwindow

func NumSubarraysWithSum(nums []int, goal int) int {
	return numSubarraysWithSum2(nums,goal)- numSubarraysWithSum2(nums,goal-1)
}

func numSubarraysWithSum1(nums []int, goal int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		j := i
		sum := 0
		for j < len(nums) && sum <= goal {
			sum += nums[j]
			if sum == goal {
				res++
			}
			j++
		}
	}
	return res
}

func numSubarraysWithSum2(nums []int, goal int) int {
	res := 0
	j := 0
	sum := 0
	if goal<0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		sum = nums[i] + sum
		for j < len(nums) && sum > goal {
			sum = sum - nums[j]
			j++
		}
		res = res + i-j + 1
	}
	return res
}

package leet

/*

fmt.Println(leet.LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
fmt.Println(leet.LongestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
*/

func LongestOnes(nums []int, k int) int {
	start := 0
	cntLen := 0
	cntZero := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		cntLen++
		if nums[i] == 0 {
			cntZero++
		}
		if res < cntLen {
			res = cntLen
		}
		for cntZero > k {
			if nums[start] == 0 {
				cntZero--
			}
			cntLen--
			start++
		}

	}
	return res - 1
}

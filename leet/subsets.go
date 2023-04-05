package leet

import (
	"fmt"
	"math"
)

func Subsets(nums []int) [][]int {
	//sort.Ints(nums)
	size := int(math.Pow(float64(2), float64(len(nums))))
	bitPos := make([][]int, size)
	for i := 0; i < len(nums); i++ {

		for j := 0; j < size; j++ {
			if (j>>i)&1 == 1 {
				if bitPos[j] == nil {
					bitPos[j] = []int{}
				}
				fmt.Println((j >> 1), j, i)
				bitPos[j] = append(bitPos[j], nums[i])
			}
		}

	}
	return bitPos
}

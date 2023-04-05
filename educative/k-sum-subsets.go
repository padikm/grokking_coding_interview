package educative

import "fmt"

func FindKSumSubsets(n []int, k int) {
	findKSumSubsets(n, k, []int{}, 0)
}

func findKSumSubsets(n []int, k int, res []int, sum int) {
	if len(n) == 0 {
		if sum == k {
			fmt.Println(res, sum)
		}

		return
	}
	findKSumSubsets(n[1:len(n)], k, append(res, n[0]), sum+n[0])
	findKSumSubsets(n[1:len(n)], k, res, sum)

}

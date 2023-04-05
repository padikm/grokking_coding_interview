package leet

func MinDeletion(nums []int) int {
	return minDeletion(nums, 0, 0)
}

func minDeletion(a []int, i int, count int) int {
	if i == len(a) || len(a) == 0 {
		return count
	}

	if i%2 != 0 && a[i] == a[i+1] {
		a = a[1:]
		return minDeletion(a, 0, count+1)
	}
	return minDeletion(a, i+1, count)
}

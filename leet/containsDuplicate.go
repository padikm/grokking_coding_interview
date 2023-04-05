package leet

func ContainsDuplicate(nums []int) bool {
	var dup = make(map[int]int)
	for _, j := range nums {
		if _, ok := dup[j]; ok {
			return true
		}
		dup[j]++
	}
	return false
}

//https://leetcode.com/problems/contains-duplicate-ii/
func ContainsNearbyDuplicate(a []int, k int) bool {
	s := make(map[int]int)
	for i := 0; i < len(a); i++ {
		s[a[i]] = -1
	}
	for i := 0; i < len(a); i++ {
		if s[a[i]] > -1 && i-s[a[i]] <= k {
			return true
		}
		s[a[i]] = i
	}
	return false
}

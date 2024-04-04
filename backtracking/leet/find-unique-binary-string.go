package leet

var s = []string{"0", "1"}

func FindDifferentBinaryString(nums []string) string {
	return findDifferentBinaryString1(nums, "")
}

func findDifferentBinaryString1(nums []string, res string) string {
	if len(res) == len(nums[0]) {
		for i := 0; i < len(nums); i++ {
			if nums[i] == res {
				return ""
			}
		}
		return res
	}
	for i := 0; i < len(s); i++ {
		res += s[i]
		r := findDifferentBinaryString1(nums, res)
		if len(r) != 0 {
			return r
		}
		res = res[:len(res)-1]
	}
	return ""
}

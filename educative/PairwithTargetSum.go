package educative

func PairWithTargetSum(n []int, t int) []int {
	start := 0
	end := len(n) - 1

	for end >= start {
		sum := n[start] + n[end]
		if sum > t {
			end--
		} else if sum < t {
			start++
		} else {
			return []int{start + 1, end + 1}
		}
	}
	return []int{}
}

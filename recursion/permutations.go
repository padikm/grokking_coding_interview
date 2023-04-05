package recursion

func PermutationRec(s string) []string {
	res := permutation(s, 0)
	return res
}

func CountPermutations(s string) int {
	return countPermutation(s, 0)
}
func permutation(s string, cur int) []string {
	res := make([]string, 0)
	if cur == len(s) {
		return []string{s}

	}
	for i := cur; i < len(s); i++ {
		r := []rune(s)
		r[cur], r[i] = r[i], r[cur]
		s = string(r)
		res = append(res, permutation(s, cur+1)...)
	}
	return res

}

func countPermutation(s string, cur int) int {
	var res int
	if cur == len(s) {
		return 1

	}
	for i := cur; i < len(s); i++ {
		r := []rune(s)
		r[cur], r[i] = r[i], r[cur]
		s = string(r)
		res = res + countPermutation(s, cur+1)
	}
	return res

}

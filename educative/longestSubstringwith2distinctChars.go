// You can edit this code!
// Click here and start typing.
package educative

func LongestSubstringWith2DistChar(s string) int {
	count := 0
	res := 0
	start := 0
	m := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		m[r]++
		count++
		for len(m) > 2 {
			if count > res {
				res = count - 1
			}
			count--
			rs := rune(s[start])
			start++
			delete(m, rs)
		}
	}
	if count > res {
		return count
	}
	return res
}

package leet

func GenerateParenthesis(n int) []string {
	res := make([]string, 0)
	generateParenthesis1(n, "", 0, 0, &res)
	return res
}

func generateParenthesis1(n int, res string, open int, closed int, r *[]string) {
	if open >= n && closed >= n {
		*r = append(*r, res)
		return
	}
	if open < n {
		generateParenthesis1(n, res+"(", open+1, closed, r)
	}
	if closed < open {
		generateParenthesis1(n, res+")", open, closed+1, r)
	}

}

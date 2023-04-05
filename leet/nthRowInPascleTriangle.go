package leet

func Generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		//nc0
		prev := 1
		res[i] = make([]int, 0)
		res[i] = append(res[i], prev)
		for j := 1; j <= i; j++ {
			prev = prev * (i - j + 1) / j
			res[i] = append(res[i], prev)
		}
	}
	return res
}

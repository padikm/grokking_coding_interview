package leet

func MatrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	var arr1D []int
	var res = make([][]int, r)
	if m*n != r*c {
		return mat
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr1D = append(arr1D, mat[i][j])
		}
	}
	k := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res[i] = append(res[i], arr1D[k])
			k++
		}
	}
	return res
}

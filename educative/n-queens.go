package educative

func SolveNQueens(n int) int {
	count := 0
	NQueensUtil([][]int{}, 0, n, &count)
	return count
}

func NQueensUtil(res [][]int, j, n int, count *int) int {
	if j == n {
		*count = *count + 1
		//fmt.Println(*count)
		//fmt.Println(res)
		return 1
	}

	for i := 0; i < n; i++ {
		//fmt.Println(i, j)
		if len(res) == 0 {
			res = make([][]int, n)
			for i := range res {
				res[i] = make([]int, n)
			}
		}

		if isSafe(res, i, j, n) {
			res[i][j] = 1
			NQueensUtil(res, j+1, n, count)
			res[i][j] = 0
		}
	}
	return *count
}

func isSafe(output [][]int, i, j, n int) bool {
	//see if row is clean
	k := 0
	for i >= 0 && k < n {
		if output[i][k] == 1 {
			return false
		}
		k++
	}
	// see if column is clean
	k = 0
	for j >= 0 && k < n {
		if output[k][j] == 1 {
			return false
		}
		k++
	}

	//upper right diagonal check
	r := i - 1
	c := j + 1
	for r >= 0 && c < n {
		if output[r][c] == 1 {
			return false
		}
		r--
		c++
	}

	//upper left diagonal check
	r = i - 1
	c = j - 1
	for r >= 0 && c >= 0 {
		if output[r][c] == 1 {
			return false
		}
		r--
		c--
	}

	//lower right diagonal check
	r = i + 1
	c = j + 1
	for r < n && c < n {
		if output[r][c] == 1 {
			return false
		}
		r++
		c++
	}

	//lower left diagonal check
	r = i + 1
	c = j - 1
	for r < n && c >= 0 {
		if output[r][c] == 1 {
			return false
		}
		r++
		c--
	}
	return true
}

/*
100
000
000
*/

package leet

func IsValidSudoku(board [][]string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			if string(board[i][j]) != "." && !isValid(board, i, j) {
				return false
			}
		}
	}
	return true
}

func isValid(board [][]string, i, j int) bool {
	//row validator
	m := make(map[int]int)
	m[0] = 0
	m[1] = 3
	m[2] = 6

	row := i / 3
	col := j / 3
	for k := 0; k < len(board); k++ {
		if string(board[i][k]) == "." {
			continue
		}
		if k != j && string(board[i][k]) == string(board[i][j]) {
			return false
		}
	}
	//column
	for k := 0; k < len(board[i]); k++ {
		if string(board[k][j]) == "." {
			continue
		}
		if k != i && string(board[k][j]) == string(board[i][j]) {
			return false
		}
	}

	row = m[row]
	col = m[col]

	for k := row; k < row+3; k++ {
		for l := col; l < col+3; l++ {
			if string(board[k][l]) == "." {
				continue
			}
			if k != i && l != j && string(board[k][l]) == string(board[i][j]) {
				return false
			}
		}
	}
	return true
}

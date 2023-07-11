package leet

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if !isValid(board, i, j) {
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, i, j int) bool {
	//row validator
	for k := 0; k < len(board); k++ {
		if k != j && string(board[i][k]) == string(board[i][j]) {
			return false
		}
	}
	//column
	for k := 0; k < len(board[i]); k++ {
		if k != i && string(board[k][j]) == string(board[i][j]) {
			return false
		}
	}

}
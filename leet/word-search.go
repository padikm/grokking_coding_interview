package leet

func Exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfsWordSearch(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func dfsWordSearch(board [][]byte, word string, row int, col int) bool {
	if len(word) == 0 {
		return true
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || string(board[row][col]) != string(word[0]) {
		return false
	}
	//back tracking
	board[row][col] = byte('*')
	for _, v := range [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		res := dfsWordSearch(board, word[1:], row+v[0], col+v[1])
		if res {
			return res
		}
	}
	//back tracking
	board[row][col] = word[0]
	return false
}

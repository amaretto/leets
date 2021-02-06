package main

func isValidSudoku(board [][]byte) bool {
	var marray [9]map[byte]bool
	var m2darray [3][3]map[byte]bool
	// search row
	for i := 0; i < 9; i++ {
		m := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if marray[j] == nil {
				marray[j] = make(map[byte]bool)
			}
			if m2darray[i/3][j/3] == nil {
				m2darray[i/3][j/3] = make(map[byte]bool)
			}
			if m[board[i][j]] || marray[j][board[i][j]] || m2darray[i/3][j/3][board[i][j]] {
				return false
			}
			m[board[i][j]] = true
			marray[j][board[i][j]] = true
			m2darray[i/3][j/3][board[i][j]] = true
		}
	}
	return true
}

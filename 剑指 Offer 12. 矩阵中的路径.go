func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {

				if search(board, i, j, 0, word) {
					return true
				}

			}
		}
	}
	return false
}
func search(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}

	if word[k] == board[i][j] {
		temp := board[i][j]
		board[i][j] = ' '
		//下面这个if语句，如果成功进入，说明找到该字符，然后进行下一个字符的搜索,直到所有的搜索都成功，
		//即k == len(word) - 1 的大小时，会返回true，进入该条件语句，然后返回函数true值。
		if search(board, i, j+1, k+1, word) ||
			search(board, i, j-1, k+1, word) ||
			search(board, i+1, j, k+1, word) ||
			search(board, i-1, j, k+1, word) {
			return true
		} else {
			//没有找到目标字符，还原之前重新赋值的board[i][j]
			board[i][j] = temp
		}
	}

	return false
}
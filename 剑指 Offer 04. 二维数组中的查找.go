func findNumberIn2DArray(matrix [][]int, target int) bool {

	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	if n == 0 {
		return false
	}
	var px int
	for i := 0; i < m; i++ {
		if matrix[i][n/2] == target {
			return true
		}
		if matrix[i][n/2] > target {
			px = i
			break
		} else {
			px = i
		}
	}
	for i := px; i >= 0; i-- {
		for j := n / 2; j < n; j++ {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				break
			}
		}
	}
	for i := px; i < m; i++ {
		for j := n / 2; j >= 0; j-- {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] < target {
				break
			}
		}

	}
	return false
}
func searchMatrix(matrix [][]int, target int) bool {
	l_i := 0
	r_i := len(matrix) - 1
	mid := 0
	for l_i <= r_i {
		mid = (l_i + r_i) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			r_i = mid - 1
		} else {
			l_i = mid + 1
		}
	}
	if matrix[mid][0] > target && mid - 1 >= 0 {
		mid -= 1
	}
	l_j := 0
	r_j := len(matrix[mid]) - 1
	mid_row := 0
	for l_j <= r_j {
		mid_row = (l_j + r_j) / 2
		if matrix[mid][mid_row] == target {
			return true
		} else if matrix[mid][mid_row] > target {
			r_j = mid_row - 1
		} else {
			l_j = mid_row + 1
		}
	}
	return false
}

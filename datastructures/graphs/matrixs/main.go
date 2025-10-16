package main

func rotateTheBox(box [][]byte) [][]byte {
	rows := len(box)
	cols := len(box[0])

	for r := range rows {
		// pointer i is positioned in the last position of one of the rows
		i := cols - 1
		for c := cols - 1; c >= 0; c-- {
			switch box[r][c] {
			case '#':
				// if theres a stone on the column then we can swap it out with ith position
				box[r][c], box[r][i] = box[r][i], box[r][c]
				// shift i to the left if the swap happens
				i--
			case '*':
				// set it to the current column - 1 because we cant place stones here
				i = c - 1
			}
		}
	}

	// rotateTheBox after performing the above operations
	res := make([][]byte, cols)
	for c := range cols {
		newRow := make([]byte, rows) // this is a row after rotating the box
		for r := range rows {
			newRow[r] = box[rows-1-r][c]
		}
		res[c] = newRow
	}
	return res
}

func rotate(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	res := make([][]int, cols)
	for c := range cols {
		newRow := make([]int, rows)
		for r := range rows {
			newRow[r] = matrix[rows-1-r][c]
		}
		res[c] = newRow
	}

	copy(matrix, res)
}

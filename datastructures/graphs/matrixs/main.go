package main

import (
	"container/list"
	"fmt"
)

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

type coords struct {
	x int
	y int
}

// dfs way is to use recursion
func uniquePathsdfs(matrix [][]int, r int, c int, visited map[coords]bool) int {
	rows, cols := len(matrix), len(matrix[0])
	if (min(r, c) < 0 || r == rows || c == cols || visited[coords{r, c}] || matrix[r][c] == 1) {
		return 0
	}
	if r == rows-1 && c == cols-1 {
		return 1
	}
	visited[coords{r, c}] = true

	count := 0
	count += uniquePathsdfs(matrix, r+1, c, visited)
	count += uniquePathsdfs(matrix, r-1, c, visited)
	count += uniquePathsdfs(matrix, r, c+1, visited)
	count += uniquePathsdfs(matrix, r, c-1, visited)

	visited[coords{r, c}] = false
	return count
}

// bfs uses a queue and great to find the shortest path
func uniquePathsbfs(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	visited := make(map[coords]bool)

	queue := list.New()
	queue.PushBack(coords{0, 0})
	visited[coords{0, 0}] = true

	length := 0
	for queue.Len() > 0 {
		queueLen := queue.Len()
		for i := 0; i < queueLen; i++ {
			pair := queue.Remove(queue.Front()).(coords)
			r, c := pair.x, pair.y

			if r == rows-1 && c == cols-1 {
				return length
			}

			neighbors := []coords{
				{r, c + 1},
				{r, c - 1},
				{r + 1, c},
				{r - 1, c},
			}
			for _, n := range neighbors {
				if min(n.x, n.y) < 0 || n.x == rows || n.y == cols || visited[n] || matrix[n.x][n.y] == 1 {
					continue
				}
				queue.PushBack(n)
				visited[n] = true
			}
		}
		length++
	}
	return -1
}

var diagonalSteps = []coord{
	{1, 1},   // ↘
	{1, -1},  // ↙
	{-1, 1},  // ↗
	{-1, -1}, // ↖
}

type coord struct {
	row int
	col int
}

func longestDiagonal(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	visited := make(map[coord]bool)
	maxPathLength := 0

	var exploreDiagonal func(row, col, length int)
	exploreDiagonal = func(row, col, length int) {
		if length > maxPathLength {
			maxPathLength = length
		}

		for _, step := range diagonalSteps {
			next := coord{row + step.row, col + step.col}

			if next.row >= 0 && next.col >= 0 &&
				next.row < rows && next.col < cols &&
				!visited[next] && matrix[next.row][next.col] == 0 {

				visited[next] = true
				exploreDiagonal(next.row, next.col, length+1)
				delete(visited, next) // backtrack
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				start := coord{r, c}
				visited[start] = true
				exploreDiagonal(r, c, 1)
				delete(visited, start)
			}
		}
	}

	return maxPathLength
}

func main() {
	matrix := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 0, 0},
	}
	fmt.Println(longestDiagonal(matrix))
}

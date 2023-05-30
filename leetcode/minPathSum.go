package leetcode

import "math"

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

// Note: You can only move either down or right at any point in time.
func MinPathSum() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			//Bound check
			var top int
			var bottom int
			var min int
			if i-1 < 0 {
				top = math.MaxInt
			} else {
				top = grid[i-1][j]
			}
			if j-1 < 0 {
				bottom = math.MaxInt
			} else {
				bottom = grid[i][j-1]
			}
			if top == math.MaxInt && bottom == math.MaxInt {
				top, bottom = 0, 0
			} else {
				min = MinVal(top, bottom)
			}
			grid[i][j] += min

		}
	}
}

func MinVal(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

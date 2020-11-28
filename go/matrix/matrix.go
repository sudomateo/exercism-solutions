package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a matrix of integers.
type Matrix [][]int

// New creates a new matrix from a given string.
func New(s string) (Matrix, error) {
	var m Matrix
	rows := strings.Split(s, "\n")
	for rowNum, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")
		if rowNum > 0 && len(m[rowNum-1]) != len(cols) {
			return nil, fmt.Errorf("uneven rows %d and %d", rowNum-1, rowNum)
		}
		nums := make([]int, 0, len(cols))
		for _, numStr := range cols {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse %s as a number: %v", numStr, err)
			}
			nums = append(nums, num)
		}
		m = append(m, nums)
	}
	return m, nil
}

// Set assigns val to the row and col specified.
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[row]) {
		return false
	}
	m[row][col] = val
	return true
}

// Rows returns a nested slice of rows.
func (m Matrix) Rows() [][]int {
	tmp := make(Matrix, len(m))
	for rowNum, row := range m {
		tmp[rowNum] = make([]int, len(row))
		copy(tmp[rowNum], row)
	}
	return tmp
}

// Cols returns a nested slice of cols.
func (m Matrix) Cols() [][]int {
	tmp := make(Matrix, len(m[0]))
	for row := range tmp {
		tmp[row] = make([]int, len(m))
		for col := 0; col < len(m); col++ {
			tmp[row][col] = m[col][row]
		}
	}
	return tmp
}

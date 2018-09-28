package main

import (
	"testing"
)

func TestCreateBoard(t *testing.T) {
	board := createBoard([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0})
	if len(board) != 3 {
		t.Error("Board is not 3 rows tall")
	}
	if len(board[0]) != 4 {
		t.Error("Board is not 4 columns wide")
	}
	for i, row := range board {
		for j, val := range row {
			if i == 2 && j == 3 && val != 0 {
				t.Error("Board incorrectly constructed 0 not at end")
			}
			if i*4+j+1 != val && val != 0 {
				t.Errorf("Board incorrectly constructed number %d out of place", val)
			}
		}
	}
}

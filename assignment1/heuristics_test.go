package main

import (
	"testing"
)

func TestManhattanCountsHorizontal(t *testing.T) {
	board := createBoard([]int{4, 2, 3, 1, 5, 6, 7, 0})
	cost := modifiedManhattanDistanceHeuristic(board)
	if cost != 6 {
		t.Errorf("Manhattan Heuristic should calculate 6 for board %s but got %f", board.key(), cost)
	}
}
func TestManhattanCountsVertical(t *testing.T) {
	board := createBoard([]int{1, 7, 3, 4, 5, 6, 2, 0})
	cost := modifiedManhattanDistanceHeuristic(board)
	if cost != 2 {
		t.Errorf("Manhattan Heuristic should calculate 2 for board %s but got %f", board.key(), cost)
	}
}

func TestManhattanCountsDiagonals(t *testing.T) {
	board := createBoard([]int{1, 2, 3, 13, 5, 6, 7, 8, 9, 10, 11, 12, 4, 14, 15, 0})
	cost := modifiedManhattanDistanceHeuristic(board)
	if cost != 6 {
		t.Errorf("Manhattan Heuristic should calculate 6 for board %s but got %f", board.key(), cost)
	}
}

func TestManhattanCountsMixedDifferences(t *testing.T) {
	board := createBoard([]int{1, 2, 3, 4, 12, 13, 7, 8, 9, 10, 11, 5, 6, 14, 15, 0})
	cost := modifiedManhattanDistanceHeuristic(board)
	if cost != 10 {
		t.Errorf("Manhattan Heuristic should calculate 10 for board %s but got %f", board.key(), cost)
	}
}

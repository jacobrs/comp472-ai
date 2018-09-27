package main

import (
	"math"
)

func tilesAwayHeuristic(b board) int {
	cost := 0
	for i, row := range b {
		for j, col := range row {
			positionInPuzzle := (i*4 + j) + 1
			if col != 0 {
				cost += int(math.Abs(float64(col - positionInPuzzle)))
			} else {
				cost += (len(row) * len(b)) - positionInPuzzle
			}
		}
	}
	return cost
}

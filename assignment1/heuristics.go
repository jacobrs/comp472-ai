package main

import (
	"math"
)

func cartesianDistanceHeuristic(b board, rowSize int) float64 {
	cost := 0.0
	for i, row := range b {
		for j, num := range row {
			if num != 0 {
				properRow, properCol := (num-1)/rowSize, (num-1)%rowSize

				rowDiff := Abs(properRow - i)
				colDiff := Abs(properCol - j)
				cost += math.Sqrt(float64(rowDiff*rowDiff + colDiff*colDiff))
			}
		}
	}
	return cost
}

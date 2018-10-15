package main

import (
	"math"
)

func cartesianDistanceHeuristic(b board) float64 {
	cost := 0.0
	for i, row := range b {
		for j, num := range row {
			if num != 0 {
				properRow, properCol := (num-1)/4, (num-1)%4

				rowDiff := Abs(properRow - i)
				colDiff := Abs(properCol - j)
				cost += math.Sqrt(float64(rowDiff*rowDiff + colDiff*colDiff))
			}
		}
	}
	return cost
}

func modifiedManhattanDistanceHeuristic(b board) float64 {
	cost := 0.0
	for i, row := range b {
		for j, num := range row {
			if num != 0 {
				properRow, properCol := (num-1)/4, (num-1)%4

				rowDiff := Abs(properRow - i)
				colDiff := Abs(properCol - j)

				diff := Abs(rowDiff - colDiff)

				cost += float64(diff + Min(rowDiff, colDiff))
			}
		}
	}
	return cost
}

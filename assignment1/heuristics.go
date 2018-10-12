package main

import (
	"math"
)

func cartesianDistanceHeuristic(b board) float64 {
	cost := 0.0
	for i, row := range b {
		for j, num := range row {
			if num != 0 {
				properRow, properCol := num/4, num%4-1

				cost += math.Sqrt(float64(Abs(properRow-i)*Abs(properRow-i) + Abs(properCol-j)*Abs(properCol-j)))
			}
		}
	}
	return cost
}

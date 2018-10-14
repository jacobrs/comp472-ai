package main

import (
	"fmt"
	"time"
)

// GameStatistics is a struct that keeps track of some stats
type GameStatistics struct {
	numberOfNodes            int
	effectiveBranchingFactor float64
	finalCost                int
	startTime                time.Time
	endTime                  time.Time
	timeTaken                time.Duration
}

// Init - A init function for tallying up our statistics, this executes at the beginning of the search
func (gs *GameStatistics) Init() {
	gs.startTime = time.Now()
}

// Step - A step function for tallying up our statistics, this executes on each step of the search
func (gs *GameStatistics) Step() {
	gs.numberOfNodes++
}

// End - A end function for tallying up our statistics, this executes at the end of the search
func (gs *GameStatistics) End(g *GameState) {
	gs.endTime = time.Now()
	gs.finalCost = g.cost
	gs.timeTaken = gs.endTime.Sub(gs.startTime)

	// Calculate effective branching factor

	fmt.Println(fmt.Sprintf("Num nodes: %d", gs.numberOfNodes))
	fmt.Println(fmt.Sprintf("EBF: %.5f", gs.effectiveBranchingFactor))
	fmt.Println(fmt.Sprintf("Final cost: %d", gs.finalCost))
	fmt.Println(fmt.Sprintf("Time taken: %s", gs.timeTaken.String()))
}

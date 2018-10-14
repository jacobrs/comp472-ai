package main

import (
	"container/heap"
	"fmt"
)

// GameState is a state of the board and its cost and heuristic value
type GameState struct {
	state     board
	hValue    float64
	cost      int
	index     int // The index of the gameState in the heap.
	moveMade  string
	parent    *GameState
	gameStats *GameStatistics
	depth     int
}

// GameStatePriorityQueue is a PriorityQueue for game states
type GameStatePriorityQueue []*GameState

func (q GameStatePriorityQueue) Len() int { return len(q) }

func (q GameStatePriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest value
	return q[i].hValue+float64(q[i].cost) < q[j].hValue+float64(q[j].cost)
}

func (q GameStatePriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

// Push - used to push something on the queue
func (q *GameStatePriorityQueue) Push(x interface{}) {
	n := len(*q)
	gameState := x.(*GameState)
	gameState.index = n
	*q = append(*q, gameState)
}

// Pop - Used for getting the top of the queue
func (q *GameStatePriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	gameState := old[n-1]
	gameState.index = -1 // for safety
	*q = old[0 : n-1]
	return gameState
}

// Update - modifies the priority and value of an gameState in the queue.
func (q *GameStatePriorityQueue) Update(gameState *GameState, hValue float64, cost int, priority int) {
	gameState.hValue = hValue
	gameState.cost = cost
	heap.Fix(q, gameState.index)
}

// Print - print the priority queue state in order
func (q GameStatePriorityQueue) Print() {
	fmt.Print("Priority Queue: ")
	pq := make(GameStatePriorityQueue, 0)
	heap.Init(&pq)

	for _, state := range q {
		pq.Push(state)
	}

	for pq.Len() > 0 {
		state := heap.Pop(&pq).(*GameState)
		fmt.Print(state.hValue)

		if pq.Len() > 0 {
			fmt.Print(", ")
		}
	}

	fmt.Print("\n")
}

func (g GameState) constructPath() []string {
	var path []string

	for g.parent != nil {
		path = append([]string{g.moveMade}, path...)
		g = *g.parent
	}

	path = append(path, "[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0]")

	return path
}

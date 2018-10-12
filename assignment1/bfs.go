package main

import (
	"container/heap"
)

func (g GameState) bfs(gn func(GameState) int, fn func(board) float64) []string {
	var currState *GameState
	seenStates := make(map[string]bool)

	g.cost = 0
	pq := make(GameStatePriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &g)

	for pq.Len() > 0 {
		currState = heap.Pop(&pq).(*GameState)

		if currState.hValue == 0 {
			return currState.constructPath()
		}

		stateKey := currState.state.key()
		_, seen := seenStates[stateKey]
		if !seen {
			seenStates[stateKey] = true

			for _, board := range currState.state.possibleMoves() {
				addState := &GameState{
					state:    board,
					hValue:   fn(board),
					cost:     0,
					parent:   currState,
					moveMade: currState.state.findBlankPosition().toLetter() + " " + board.key(),
				}
				heap.Push(&pq, addState)
			}
		}
	}

	return nil
}

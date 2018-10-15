package main

import (
	"container/heap"
	"strconv"
)

// Abs - Helper for returning the absolute value of an integer
func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

// Min - Helper for returning the smallest value between two integers
func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func (g GameState) genericSearch(gn func(GameState) int, fn func(board) float64) []string {
	var currState *GameState
	seenStates := make(map[string]bool)

	g.hValue = fn(g.state)
	g.cost = gn(g)
	g.depth = 1
	g.gameStats.Init()
	pq := make(GameStatePriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &g)

	for pq.Len() > 0 {
		currState = heap.Pop(&pq).(*GameState)
		g.gameStats.Step()

		if currState.hValue == 0 {
			g.gameStats.End(currState)
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
					cost:     gn(*currState),
					depth:    currState.depth + 1,
					parent:   currState,
					moveMade: board.findBlankPosition().toLetter() + " " + board.key(),
				}
				heap.Push(&pq, addState)
			}
		}
	}

	return nil
}

func parseBoard(input []string, rowSize int) board {
	numList := []int{}
	for _, i := range input {
		num, e := strconv.Atoi(i)
		if e == nil {
			numList = append(numList, num)
		}
	}
	return createBoard(numList, rowSize)
}

func reverse(arr []board) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

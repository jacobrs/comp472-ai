package main

import (
	stack "github.com/golang-collections/collections/stack"
)

// Iterative DFS algorithm
func (g GameState) dfs(goal string, maxDepth int) []string {
	stateStack := stack.New()
	stateStack.Push(&g)

	for stateStack.Len() > 0 {
		currState := stateStack.Pop().(*GameState)
		currDepth := currState.depth

		if currState.state.key() == goal {
			return currState.constructPath()
		}

		if currDepth > maxDepth && maxDepth >= 0 {
			continue
		}

		for _, board := range currState.state.possibleMoves() {
			addState := &GameState{
				state:    board,
				hValue:   0,
				cost:     0,
				depth:    currState.depth + 1,
				parent:   currState,
				moveMade: board.findBlankPosition().toLetter() + " " + board.key(),
			}

			stateStack.Push(addState)
		}
	}

	return []string{}
}

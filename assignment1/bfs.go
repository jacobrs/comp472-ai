package main

func (g GameState) bestFirstSearch(fn func(board) float64) []string {
	return g.genericSearch(func(g GameState) int { return 0 }, fn)
}

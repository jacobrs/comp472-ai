package main

func (g GameState) aSearch(fn func(board) float64) []string {
	return g.genericSearch(func(g GameState) int { return int(g.parent.cost + 1) }, fn)
}

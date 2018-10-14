package main

func (g GameState) aSearch(fn func(board) float64) []string {
	return g.genericSearch(aSearchCost, fn)
}

func aSearchCost(g GameState) int {
	if g.parent != nil {
		return int(g.parent.cost + 1)
	}
	return 0
}

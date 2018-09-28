package main

func (b board) bfs(path []string) []string {
	currentScore := tilesAwayHeuristic(b)
	if currentScore == 0 {
		return path
	}
	children := sort(b.possibleMoves())
	for _, child := range children {
		if !contains(path, child.key()) {
			childScore := tilesAwayHeuristic(child)
			if childScore <= currentScore {
				result := child.bfs(append(path, child.findBlankPosition().toLetter()+" "+child.key()))
				if len(result) > 0 {
					return result
				}
			}
		}
	}
	return []string{}
}

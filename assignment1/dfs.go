package main

func (b board) dfs(visited *[]string, path []string, goal string, maxDepth int, currentDepth int) []string {
	if b.key() == goal {
		return path
	}
	if currentDepth > maxDepth {
		return []string{}
	}
	*visited = append(*visited, b.key())
	// nvisited := len(*visited)
	// print("\033[H\033[2J")
	// fmt.Printf("%d nodes visited\n", nvisited)
	for _, child := range b.possibleMoves() {
		if !contains(*visited, child.key()) {
			result := child.dfs(visited, append(path, child.findBlankPosition().toLetter()+" "+child.key()), goal, maxDepth, currentDepth+1)
			if len(result) > 0 {
				return result
			}
		}
	}
	return []string{}
}

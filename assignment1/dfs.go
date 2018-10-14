package main

func (b board) dfs(visited *[]string, path []string, goal string, maxDepth int, currentDepth int) []string {
	if b.key() == goal {
		return path
	}
	if currentDepth > maxDepth && maxDepth >= 0 {
		return []string{}
	}
	*visited = append(*visited, b.key())
	for _, child := range b.possibleMoves() {
		if !contains(path, child.key()) {
			result := child.dfs(visited, append(path, child.findBlankPosition().toLetter()+" "+child.key()), goal, maxDepth, currentDepth+1)
			if len(result) > 0 {
				return result
			}
		}
	}
	return []string{}
}

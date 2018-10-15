package main

func (b board) dfs(visited map[string]bool, path []string, goal string, maxDepth int, currentDepth int) []string {
	if b.key() == goal {
		return path
	}
	if currentDepth > maxDepth && maxDepth >= 0 {
		return []string{}
	}
	visited[b.key()] = true
	for _, child := range b.possibleMoves() {
		if !visited[child.key()] {
			result := child.dfs(visited, append(path, child.findBlankPosition().toLetter()+" "+child.key()), goal, maxDepth, currentDepth+1)
			if len(result) > 0 {
				visited[b.key()] = false
				return result
			}
		}
	}
	visited[b.key()] = false
	return []string{}
}

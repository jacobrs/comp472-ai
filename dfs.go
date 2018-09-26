package main

import (
	"fmt"
)

func (b board) dfs(visited *[]string, path []string, goal string) []string {
	if b.key() == goal {
		fmt.Println("Goal found!")
		return path
	}
	if contains(*visited, b.key()) {
		return []string{}
	}
	*visited = append(*visited, b.key())
	for _, child := range b.possibleMoves() {
		result := child.dfs(visited, append(path, child.key()), goal)
		if len(result) > 0 {
			return result
		}
	}
	return []string{}
}

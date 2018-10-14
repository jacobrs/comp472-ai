package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	searchAlgorithm := "dfs"
	var b board

	if len(os.Args) > 1 {
		strList := (strings.Split(os.Args[1], " "))
		numList := []int{}
		for _, i := range strList {
			num, e := strconv.Atoi(i)
			if e == nil {
				numList = append(numList, num)
			}
		}
		b = createBoard(numList)

		if len(os.Args) > 2 {
			searchAlgorithm = os.Args[2]
		}
	} else {
		b = getBoard(3, 4)
	}

	if searchAlgorithm == "dfs" {
		fmt.Println("Running depth first search")
		prettyPrintPath(b.dfs(&[]string{}, []string{}, "[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0]", -1, 1))
	} else if searchAlgorithm == "bfs" {
		// Running on sequential BFS
		fmt.Println("Running best first search")
		var game GameState
		game.state = b
		game.cost = 0
		game.hValue = cartesianDistanceHeuristic(b)
		fmt.Println(game.bestFirstSearch(cartesianDistanceHeuristic))
	} else {
		// Running on sequential BFS
		fmt.Println("Running A search")
		var game GameState
		game.state = b
		game.cost = 0
		game.hValue = cartesianDistanceHeuristic(b)
		fmt.Println(game.aSearch(cartesianDistanceHeuristic))
	}
}

func prettyPrintPath(path []string) {
	for _, pos := range path {
		fmt.Println(pos)
	}
}

func contains(keys []string, key string) bool {
	for _, i := range keys {
		if i == key {
			return true
		}
	}
	return false
}

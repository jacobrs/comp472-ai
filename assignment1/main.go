package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		strList := (strings.Split(os.Args[1], " "))
		numList := []int{}
		for _, i := range strList {
			num, e := strconv.Atoi(i)
			if e == nil {
				numList = append(numList, num)
			}
		}
		b := createBoard(numList)

		// Running on sequential BFS
		fmt.Println("Running best first search with user input")
		var game GameState
		game.state = b
		game.cost = 0
		game.hValue = cartesianDistanceHeuristic(b)
		fmt.Println(game.bestFirstSearch(cartesianDistanceHeuristic))
	} else {
		getBoard(3, 4).print()
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

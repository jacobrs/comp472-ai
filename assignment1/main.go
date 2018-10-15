package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	programMode := "assignment"

	if len(os.Args) > 1 {
		programMode = os.Args[1]
	} else {
		fmt.Println("Invalid command format")
	}

	if programMode == "assignment" {
		assignmentMode()
	} else {
		experimentMode()
	}
}

func assignmentMode() {
	searchAlgorithm := "dfs"
	var b board

	if len(os.Args) > 2 {
		searchAlgorithm = os.Args[2]
	}

	if len(os.Args) > 3 {
		strList := (strings.Split(os.Args[3], " "))
		b = parseBoard(strList, 4)
	} else {
		b = getBoard(3, 4)
	}

	if searchAlgorithm == "dfs" {
		fmt.Println("Running depth first search")
		visited := make(map[string]bool)
		prettyPrintPath(b.dfs(visited, []string{}, "[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0]", -1, 1))
	} else if searchAlgorithm == "bfs" {
		// Running on sequential BFS
		fmt.Println("Running best first search")
		var game GameState
		game.state = b
		game.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
		heuristic := func(b board) float64 { return cartesianDistanceHeuristic(b, 4) }
		fmt.Println(game.bestFirstSearch(heuristic))
	} else if searchAlgorithm == "astar" {
		// Running on sequential BFS
		fmt.Println("Running A search")
		var game GameState
		game.state = b
		game.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
		heuristic := func(b board) float64 { return cartesianDistanceHeuristic(b, 4) }
		fmt.Println(game.aSearch(heuristic))
	} else {
		// Run against DFS, BFS with h1, BFS with h2, AStar with h1, and AStar with h2
	}
}

func experimentMode() {
	var b board

	if len(os.Args) > 2 && os.Args[2] == "custom" {
		_, rowSize := getBoardDimensionsFromCLI(4)
		heuristic := func(b board) float64 { return cartesianDistanceHeuristic(b, rowSize) }
		strList := (strings.Split(os.Args[3], " "))
		b = parseBoard(strList, rowSize)
		runExperiment(b, heuristic)
	} else {
		numRuns := 5
		var e error
		if len(os.Args) > 2 {
			numRuns, e = strconv.Atoi(os.Args[2])

			if e != nil {
				fmt.Println("Was expecting an integer for the number of runs")
				os.Exit(1)
			}
		}

		amtRows, rowSize := getBoardDimensionsFromCLI(3)
		heuristic := func(b board) float64 { return cartesianDistanceHeuristic(b, rowSize) }

		for i := 0; i < int(numRuns); i++ {
			// Generate a random board
			b = getBoard(amtRows, rowSize)

			runExperiment(b, heuristic)
		}
	}
}

func getBoardDimensionsFromCLI(amtRowsIndex int) (int, int) {
	var amtRows int
	var rowSize int
	if len(os.Args) > amtRowsIndex+1 {
		var e error
		amtRows, e = strconv.Atoi(os.Args[amtRowsIndex])

		if e != nil {
			amtRows = 3
		}

		rowSize, e = strconv.Atoi(os.Args[amtRowsIndex+1])

		if e != nil {
			rowSize = 4
		}
	} else {
		amtRows = 3
		rowSize = 4
	}

	return amtRows, rowSize
}

func runExperiment(b board, heuristic func(board) float64) {
	var gameBFS GameState
	gameBFS.state = b
	gameBFS.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
	fmt.Println(fmt.Sprintf("Running best first search on %s", b.key()))
	fmt.Println(gameBFS.bestFirstSearch(heuristic))

	var gameAStar GameState
	gameAStar.state = b
	gameAStar.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
	fmt.Println(fmt.Sprintf("Running A* search on %s", b.key()))
	fmt.Println(gameAStar.aSearch(heuristic))
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

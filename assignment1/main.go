package main

import (
	"fmt"
	"io/ioutil"
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
	var rowSize = 4
	var b board
	searchAlgorithm := "dfs"

	if len(os.Args) > 2 {
		searchAlgorithm = os.Args[2]
	}

	if len(os.Args) > 3 {
		strList := (strings.Split(os.Args[3], " "))
		_, rowSize = getBoardDimensionsFromCLI(4)
		b = parseBoard(strList, rowSize)
	} else {
		b = getBoard(3, 4)
	}

	b.print()

	if searchAlgorithm == "idfs" {
		iterativeDepthDFS(b)
	} else if searchAlgorithm == "dfs" {
		dfs(b)
	} else if searchAlgorithm == "bfs" {
		// Running on sequential BFS
		bfs(b, rowSize)
	} else if searchAlgorithm == "astar" {
		// Running on sequential BFS
		astar(b, rowSize)
	} else {
		// Run against IterativeDFS, DFS, BFS with h1, BFS with h2, AStar with h1, and AStar with h2
		astar(b, rowSize)
		bfs(b, rowSize)
		iterativeDepthDFS(b)
		dfs(b)
	}
}

func astar(b board, rowSize int) {
	fmt.Println("Running A search h1")
	var game GameState
	game.state = b
	game.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
	heuristic := func(b board) float64 { return cartesianDistanceHeuristic(b, rowSize) }
	output := game.aSearch(heuristic)
	prettyPrintPath(output)
	writePrettyPath(output, "./puzzleAs-h1.txt")
	fmt.Println("Running A search h2")
	game.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
	heuristic = func(b board) float64 { return modifiedManhattanDistanceHeuristic(b, rowSize) }
	output = game.aSearch(heuristic)
	prettyPrintPath(output)
	writePrettyPath(output, "./puzzleAs-h2.txt")
}

func bfs(b board, rowSize int) {
	var game GameState
	fmt.Println("Running best first search h1")
	game.state = b
	game.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
	heuristic := func(b board) float64 { return cartesianDistanceHeuristic(b, rowSize) }
	output := game.bestFirstSearch(heuristic)
	prettyPrintPath(output)
	writePrettyPath(output, "./puzzleBFS-h1.txt")
	fmt.Println("Running best first search h2")
	game.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
	heuristic = func(b board) float64 { return modifiedManhattanDistanceHeuristic(b, rowSize) }
	output = game.bestFirstSearch(heuristic)
	prettyPrintPath(output)
	writePrettyPath(output, "./puzzleBFS-h2.txt")
}

func iterativeDepthDFS(b board) {
	fmt.Println("Running iterative depth first search")
	path := []string{}
	maxDepth := 10
	for len(path) <= 0 {
		visited := make(map[string]bool)
		path = b.dfs(visited, []string{"0 " + b.key()}, b.goalState(), maxDepth, 1)
		maxDepth *= 2
	}
	prettyPrintPath(path)
	writePrettyPath(path, "./puzzleIterativeDFS.txt")
}

func dfs(b board) {
	fmt.Println("Running depth first search")
	visited := make(map[string]bool)
	path := b.dfs(visited, []string{"0 " + b.key()}, b.goalState(), -1, 1)
	prettyPrintPath(path)
	writePrettyPath(path, "./puzzleDFS.txt")
}

func experimentMode() {
	var b board

	if len(os.Args) > 2 && os.Args[2] == "custom" {
		_, rowSize := getBoardDimensionsFromCLI(4)
		fmt.Println("Running against cartesianDistanceHeuristic")
		heuristic := func(b board) float64 { return cartesianDistanceHeuristic(b, rowSize) }
		strList := (strings.Split(os.Args[3], " "))
		b = parseBoard(strList, rowSize)
		runExperiment(b, heuristic)

		fmt.Println("Running against modifiedManhattanDistanceHeuristic")
		heuristic = func(b board) float64 { return modifiedManhattanDistanceHeuristic(b, rowSize) }
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
		heuristicEuclidean := func(b board) float64 { return cartesianDistanceHeuristic(b, rowSize) }
		heuristicManhatten := func(b board) float64 { return modifiedManhattanDistanceHeuristic(b, rowSize) }

		for i := 0; i < int(numRuns); i++ {
			// Generate a random board
			b = getBoard(amtRows, rowSize)

			fmt.Println("Running against cartesianDistanceHeuristic")
			runExperiment(b, heuristicEuclidean)

			fmt.Println("Running against modifiedManhattanDistanceHeuristic")
			runExperiment(b, heuristicManhatten)
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

func writePrettyPath(path []string, filename string) {
	var data []byte
	for _, line := range path {
		data = append(data, []byte(line+"\n")...)
	}
	ioutil.WriteFile(filename, data, 0666)
}

func contains(keys []string, key string) bool {
	for _, i := range keys {
		if i == key {
			return true
		}
	}
	return false
}

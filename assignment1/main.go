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
		numList := []int{}
		for _, i := range strList {
			num, e := strconv.Atoi(i)
			if e == nil {
				numList = append(numList, num)
			}
		}
		b = createBoard(numList)
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
		game.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
		fmt.Println(game.bestFirstSearch(cartesianDistanceHeuristic))
	} else {
		// Running on sequential BFS
		fmt.Println("Running A search")
		var game GameState
		game.state = b
		game.gameStats = &GameStatistics{0, 0.0, 0, time.Time{}, time.Time{}, 0}
		fmt.Println(game.aSearch(cartesianDistanceHeuristic))
	}
}

func experimentMode() {
	var b board
	var numRuns int64 = 5
	var e error

	if len(os.Args) > 2 {
		numRuns, e = strconv.ParseInt(os.Args[2], 10, 8)

		if e != nil {
			fmt.Println("Was expecting an integer for the number of runs")
			os.Exit(1)
		}
	}

	if len(os.Args) > 3 {
		strList := (strings.Split(os.Args[3], " "))
		numList := []int{}
		for _, i := range strList {
			num, e := strconv.Atoi(i)
			if e == nil {
				numList = append(numList, num)
			}
		}
		b = createBoard(numList)
		runExperiment(b, cartesianDistanceHeuristic)
	} else {
		for i := 0; i < int(numRuns); i++ {
			b = getBoard(3, 4)

			runExperiment(b, cartesianDistanceHeuristic)
		}
	}
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

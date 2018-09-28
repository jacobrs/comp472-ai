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
		fmt.Println("Running dfs with 5 depth limit")
		fmt.Println("0 " + b.key())
		for _, pos := range b.dfs(&[]string{}, []string{}, "[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0]", 10, 1) {
			fmt.Println(pos)
		}
		fmt.Println("Running bfs with tilesAway heuristics")
		// positions := []board{createBoard([]int{1, 2, 3, 0, 5, 6, 7, 8, 9, 10, 4, 11}), createBoard([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0})}
		// positions = sort(positions)
		// fmt.Println(positions)
		fmt.Println("0 " + b.key())
		for _, pos := range b.bfs([]string{}) {
			fmt.Println(pos)
		}
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

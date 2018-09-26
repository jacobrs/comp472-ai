package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		strList := (strings.Split(os.Args[1], ","))
		numList := []int{}
		for _, i := range strList {
			num, e := strconv.Atoi(i)
			if e == nil {
				numList = append(numList, num)
			}
		}
		board := createBoard(numList)
		board.print()
		for i, pos := range board.dfs(&[]string{}, []string{}, "1-2-3-4-5-6-7-8-9-10-11-0") {
			fmt.Printf("%d. %s\n", i, pos)
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

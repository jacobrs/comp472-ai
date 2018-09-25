package main

import (
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
		createBoard(numList).print()
	} else {
		getBoard(3, 4).print()
	}
}

func dfs() {}

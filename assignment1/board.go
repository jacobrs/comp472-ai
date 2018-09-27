package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type board [][]int

type position struct {
	row int
	col int
}

func getBoard(n int, m int) board {
	seed := rand.NewSource(time.Now().UnixNano())
	list := rand.New(seed).Perm(n * m)
	return createBoard(list)
}

func createBoard(list []int) board {
	board := board{}
	rowSize := 4
	for i := 0; i < len(list); i++ {
		row := []int{}
		for x := 0; x < rowSize; x++ {
			row = append(row, list[0])
			list = list[1:]
		}
		board = append(board, row)
	}
	return board
}

func createBoardFromKey(key string) board {
	formattedKey := key[3 : len(key)-1]
	strList := (strings.Split(formattedKey, ","))
	numList := []int{}
	for _, i := range strList {
		num, e := strconv.Atoi(strings.Replace(i, " ", "", -1))
		if e == nil {
			numList = append(numList, num)
		}
	}
	return createBoard(numList)
}

func (b board) print() {
	fmt.Printf("Key: %s\n", b.key())
	fmt.Println("---------------------")
	for _, row := range b {
		fmt.Print("|")
		for _, col := range row {
			if col < 10 {
				fmt.Print(" ")
			}
			fmt.Printf(" %d |", col)
		}
		fmt.Println("\n---------------------")
	}
}

func (b board) key() string {
	key := "["
	for _, row := range b {
		for _, val := range row {
			key += strconv.Itoa(val) + ", "
		}
	}
	return key[:len(key)-2] + "]"
}

func (b board) possibleMoves() []board {
	positionOfBlank := b.findBlankPosition()
	possibilities := []board{}
	if positionOfBlank.row < len(b)-1 {
		// can move down
		possibilities = append(possibilities, b.swap(position{positionOfBlank.row, positionOfBlank.col},
			position{positionOfBlank.row + 1, positionOfBlank.col}))
		if positionOfBlank.col > 0 {
			// can move down left
			possibilities = append(possibilities, b.swap(position{positionOfBlank.row, positionOfBlank.col},
				position{positionOfBlank.row + 1, positionOfBlank.col - 1}))
		}
	}
	if positionOfBlank.col > 0 {
		// can move left
		possibilities = append(possibilities, b.swap(position{positionOfBlank.row, positionOfBlank.col},
			position{positionOfBlank.row, positionOfBlank.col - 1}))
	}
	if positionOfBlank.row > 0 {
		if positionOfBlank.col > 0 {
			// can move up left
			possibilities = append(possibilities, b.swap(position{positionOfBlank.row, positionOfBlank.col},
				position{positionOfBlank.row - 1, positionOfBlank.col - 1}))
		}
		// can move up
		possibilities = append(possibilities, b.swap(position{positionOfBlank.row, positionOfBlank.col},
			position{positionOfBlank.row - 1, positionOfBlank.col}))
		if positionOfBlank.col < len(b[0])-1 {
			// can move up right
			possibilities = append(possibilities, b.swap(position{positionOfBlank.row, positionOfBlank.col},
				position{positionOfBlank.row - 1, positionOfBlank.col + 1}))
		}
	}
	if positionOfBlank.col < len(b[0])-1 {
		// can move right
		possibilities = append(possibilities, b.swap(position{positionOfBlank.row, positionOfBlank.col},
			position{positionOfBlank.row, positionOfBlank.col + 1}))
	}
	if positionOfBlank.row < len(b)-1 {
		if positionOfBlank.col < len(b[0])-1 {
			// can move down right
			possibilities = append(possibilities, b.swap(position{positionOfBlank.row, positionOfBlank.col},
				position{positionOfBlank.row + 1, positionOfBlank.col + 1}))
		}
	}
	return possibilities
}

func (b board) findBlankPosition() position {
	return b.findPosition(0)
}

func (b board) findPosition(target int) position {
	for i, row := range b {
		for j, val := range row {
			if val == target {
				return position{i, j}
			}
		}
	}
	return position{0, 0}
}

func (b board) swap(p1 position, p2 position) board {
	nb := [][]int{}
	for i, row := range b {
		nb = append(nb, make([]int, len(row)))
		for x, col := range row {
			nb[i][x] = col
		}
	}
	tmp := nb[p1.row][p1.col]
	nb[p1.row][p1.col] = nb[p2.row][p2.col]
	nb[p2.row][p2.col] = tmp
	return nb
}

func (p position) toLetter() string {
	return string((p.row * 4) + p.col + 97)
}

func sort(positions []board) []board {
	boards := positions
	sorted := []board{}
	for len(boards) > 0 {
		idx := minTilesAway(boards)
		sorted = append(sorted, boards[idx])
		boards = boards[idx:len(boards)]
	}
	return positions
}

func minTilesAway(positions []board) int {
	minPosition := 0
	min := tilesAwayHeuristic(positions[minPosition])
	for idx, val := range positions {
		cost := tilesAwayHeuristic(val)
		if cost < min {
			minPosition = idx
			min = cost
		}
	}
	return minPosition
}

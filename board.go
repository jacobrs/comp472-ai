package main

import (
	"fmt"
	"math/rand"
	"time"
)

type board [][]int

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

func (b board) print() {
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

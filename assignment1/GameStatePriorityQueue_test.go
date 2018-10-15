package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"testing"
)

func TestCreateProrityQueue(t *testing.T) {
	pq := make(GameStatePriorityQueue, 0)
	heap.Init(&pq)

	g := &GameState{
		state:    nil,
		hValue:   0,
		cost:     0,
		index:    0,
		parent:   nil,
		moveMade: "",
	}

	heap.Push(&pq, g)

	if pq.Len() != 1 {
		t.Error("PriorityQueue was not properly initialized")
	}
}

func TestCreateProrityQueueOrder(t *testing.T) {
	pq := make(GameStatePriorityQueue, 0)
	heap.Init(&pq)

	g := &GameState{
		state:    nil,
		hValue:   0.0,
		cost:     0.0,
		index:    0,
		parent:   nil,
		moveMade: "",
	}

	r := rand.New(rand.NewSource(95))
	for i := 0; i < 100; i++ {
		currState := *g
		currState.hValue = r.Float64() * 1000
		heap.Push(&pq, &currState)
	}

	prevState := heap.Pop(&pq).(*GameState)
	for pq.Len() > 0 {
		nextState := heap.Pop(&pq).(*GameState)
		pq.Print()
		if nextState.hValue < prevState.hValue {
			t.Error(fmt.Sprintf("Something went wrong Misorder: %f.3 and %f.3", prevState.hValue, nextState.hValue))
		}
		prevState = nextState
	}
}

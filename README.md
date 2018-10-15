# AI Assignments
### COMP 472 fall semester 2018

### About this repository
This repository stores mini projects that have been completed during the course of the class. This currently includes a state space search mini project with a modified 8 puzzle game.

### Directory Structure

```
root
│    README.md
└─── assignment1
    │    main.go
    │    board.go
    │    asearch.go
    │    board_test.go
    │    dfs.go
    │    heuristics.go
    │    GameStatePriorityQueue.go
    │    GameStatePriorityQueue_test.go
    │    helpers.go
    └─── bfs.go
```

### CLI

## CLI Names

```
searchType = "dfs" or "bfs" or "astar", for any other string it runs the full assignment flow
boardInput = "[space separated numbers representing the board's initial state]"
numRuns = number of runs you want to run in experiment mode
numRows and numCols = need to be specified in experiment mode if you switch the board size
```

## Using assignment mode

The currently supported command for experiment mode takes the following input format:

`./builtBinary "assignment" [searchType] [boardInput]`

## Using experiment mode

The currently supported command for experiment mode takes the following input format:

`./builtBinary "experiment" [[numRuns [numRows numCols]], ["custom" boardInput numRows numCols]]`

Where a `space` within `[]` represents an AND statement and a `,` within within `[]` represents an OR statement and `[]` represents it being optional.
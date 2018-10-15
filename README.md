# AI Assignments
### COMP 472 fall semester 2018

### Step-by-step instructions
If you follow the following steps you will get the output of the first assignment, the rest of the readme explains how to use more options that our program provides:

```
For a random puzzle:
    1. Change directory to the assignment1 directory
    2. Run `go build` to build the assignment
    3. Run `./assignment1 "assignment"`, where assignment1 is the built binary file

For a specified puzzle:
    1. Change directory to the assignment1 directory
    2. Run `go build` to build the assignment
    3. Run `./assignment1 "assignment" "all" "[space separated puzzle]"`, where assignment1 is the built binary file
    4. Example: `./assignment1 "assignment" "all" "1 0 3 7 5 2 6 4 9 10 11 8"`
```

### About this repository
This repository stores mini projects that have been completed during the course of the class. This currently includes a state space search mini project with a modified 8 puzzle game.

### Definition of a Move
For this project we chose to use the definition of a move as moving INTO the blank tile, that is if a tile moves UP into the blank spot then it is considered to be an UP move.

### Directory Structure

```
root
│    README.md
└─── assignment1
    │    main.go
    │    bfs.go
    │    board.go
    │    board_test.go
    │    asearch.go
    │    dfs.go
    │    GameStatePriorityQueue.go
    │    GameStatePriorityQueue_test.go
    │    helpers.go
    │    heuristics.go
    │    heuristics_test.go
    └─── main.go
```

### CLI Names

```
searchType = "dfs" or "bfs" or "astar", for any other string it runs the full assignment flow
boardInput = "[space separated numbers representing the board's initial state]"
numRuns = number of runs you want to run in experiment mode
numRows and numCols = need to be specified in experiment mode if you switch the board size
```

### Using assignment mode

The currently supported command for experiment mode takes the following input format:

`./builtBinary "assignment" [searchType] [boardInput [numRows numCols]]`

Examples:

```
./builtBinary "assignment"                                       // Run all algorithms against a random board
./builtBinary "assignment" "dfs"                                 // Run dfs against a random board
./builtBinary "assignment" "all" "11 10 9 8 7 6 5 4 3 2 1 0"     // Run all algorithms against the specified board, default size is 3 by 4
./builtBinary "assignment" "all" "11 10 9 8 7 6 5 4 3 2 1 0" 2 6 // Run all algorithms against the specified board, given the specified dimensions
```

### Using experiment mode

The currently supported command for experiment mode takes the following input format:

`./builtBinary "experiment" [[numRuns [numRows numCols]], ["custom" boardInput [numRows numCols]]]`

Examples:

```
./builtBinary "experiment"                                          // Run 5 experiments against random boards
./builtBinary "experiment" 10                                       // Run 10 experiments against random boards
./builtBinary "experiment" 10 5 5                                   // Run 10 experiments against a random 5 by 5 board
./builtBinary "experiment" "custom" "11 10 9 8 7 6 5 4 3 2 1 0"     // Run experiments against specified board, default size is 3 by 4
./builtBinary "experiment" "custom" "11 10 9 8 7 6 5 4 3 2 1 0" 5 5 // Run experiments against specified board and size
```

Where a `space` within `[]` represents an AND statement and a `,` within within `[]` represents an OR statement and `[]` represents it being optional.

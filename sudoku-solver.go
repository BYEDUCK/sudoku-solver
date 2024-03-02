package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	RECURSIVE_CALL_LIMIT = 1_000_000
)

type SudokuSolver struct {
	rowsOccupiedNumbers    [][]bool
	columnsOccupiedNumbers [][]bool
	squaresOccupiedNumbers [][]bool
	emptyRows              []int
	emptyColumns           []int
	StartingBoard          [][]int
	SolutionBoard          [][]int
	recursiveCallsCount    int
	Solved                 bool
	stopped                bool
	S                      int
	n                      int
	n_2                    int
	SolvingTime            time.Duration
}

func NewSolver(s int) *SudokuSolver {
	solver := &SudokuSolver{}
	solver.S = s
	solver.n = s * s
	solver.n_2 = solver.n * solver.n
	solver.rowsOccupiedNumbers = make([][]bool, solver.n)
	solver.columnsOccupiedNumbers = make([][]bool, solver.n)
	solver.squaresOccupiedNumbers = make([][]bool, solver.n)
	for i := 0; i < solver.n; i++ {
		solver.rowsOccupiedNumbers[i] = make([]bool, solver.n+1)
		solver.columnsOccupiedNumbers[i] = make([]bool, solver.n+1)
		solver.squaresOccupiedNumbers[i] = make([]bool, solver.n+1)
	}
	solver.emptyRows = make([]int, solver.n_2)
	solver.emptyColumns = make([]int, solver.n_2)
	return solver
}

func (s *SudokuSolver) init(board [][]int) {
	s.recursiveCallsCount = 0
	s.stopped = false
	s.Solved = false
	s.StartingBoard = board
	s.SolutionBoard = board
	for r := 0; r < s.n; r++ {
		for c := 0; c < s.n; c++ {
			if s.StartingBoard[r][c] != 0 {
				s.rowsOccupiedNumbers[r][s.StartingBoard[r][c]] = true
				s.columnsOccupiedNumbers[c][s.StartingBoard[r][c]] = true
				s.squaresOccupiedNumbers[s.getSquareIndex(r, c)][s.StartingBoard[r][c]] = true
			}
		}
	}
	i := 0
	for r := 0; r < s.n; r++ {
		for c := 0; c < s.n; c++ {
			if s.StartingBoard[r][c] == 0 {
				s.emptyRows[i] = r
				s.emptyColumns[i] = c
				i += 1
			}
		}
	}
	for ; i < s.n_2; i++ {
		s.emptyRows[i] = -1
		s.emptyColumns[i] = -1
	}
}

func (s *SudokuSolver) Solve(board [][]int) {
	s.init(board)
	start := time.Now()
	s.solveRecursive(0, &s.SolutionBoard)
	s.SolvingTime = time.Since(start)
}

func (s *SudokuSolver) solveRecursive(currentEmptyIdx int, board *[][]int) {
	if s.Solved || s.stopped {
		return
	}
	if currentEmptyIdx >= s.n_2 {
		s.Solved = true
		return
	}
	if s.recursiveCallsCount > RECURSIVE_CALL_LIMIT {
		fmt.Println("Recursive calls limit reached")
		s.stopped = true
		return
	}
	s.recursiveCallsCount++
	r := s.emptyRows[currentEmptyIdx]
	c := s.emptyColumns[currentEmptyIdx]
	if r < 0 && c < 0 {
		s.Solved = true
		return
	}
	possibilities := s.getNumberOptionsShuffled()
	for _, p := range possibilities {
		if s.isNumberGood(r, c, p) {
			(*board)[r][c] = p
			s.rowsOccupiedNumbers[r][p] = true
			s.columnsOccupiedNumbers[c][p] = true
			s.squaresOccupiedNumbers[s.getSquareIndex(r, c)][p] = true
			s.solveRecursive(currentEmptyIdx+1, board)
			if !s.Solved {
				(*board)[r][c] = 0
				s.rowsOccupiedNumbers[r][p] = false
				s.columnsOccupiedNumbers[c][p] = false
				s.squaresOccupiedNumbers[s.getSquareIndex(r, c)][p] = false
			}
		}
	}
}

func (s *SudokuSolver) isNumberGood(r int, c int, num int) bool {
	return !(s.rowsOccupiedNumbers[r][num] ||
		s.columnsOccupiedNumbers[c][num] ||
		s.squaresOccupiedNumbers[s.getSquareIndex(r, c)][num])
}

func (s *SudokuSolver) getNumberOptionsShuffled() []int {
	arr := make([]int, s.n)
	for i := 0; i < s.n; i++ {
		arr[i] = i + 1
	}
	rand.Shuffle(len(arr), func(i int, j int) {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	})
	return arr
}

func (s *SudokuSolver) getSquareIndex(r int, c int) int {
	return r/s.S*s.S + c/s.S
}

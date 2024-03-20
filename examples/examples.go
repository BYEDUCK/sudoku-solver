package examples

import (
	"fmt"
	"strconv"
	"strings"

	sudokusolver "github.com/byeduck/sudoku-solver"
)

func prettyPrint(board [][]int, squareSize int, cellWidth int) {
	rowLength := squareSize*squareSize*cellWidth + (squareSize + 1)
	rowSeparator := strings.Repeat("-", rowLength)
	for rowIdx, row := range board {
		if rowIdx%squareSize == 0 {
			fmt.Println(rowSeparator)
		}
		for cellIdx, cell := range row {
			if cellIdx%squareSize == 0 {
				fmt.Print("|")
			}
			cellValue := strconv.FormatInt(int64(cell), 10)
			if len(cellValue) < cellWidth {
				cellValue = cellValue + strings.Repeat(" ", cellWidth-len(cellValue))
			}
			fmt.Print(cellValue)
		}
		fmt.Print("|\n")
	}
	fmt.Println(rowSeparator)
}

func Example1() {
	var board [][]int = [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 6, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
	}
	solver3 := sudokusolver.NewSolver(3)
	solver3.Solve(board)
	fmt.Printf("Sudoku solved in %v\n", solver3.SolvingTime)
	prettyPrint(solver3.SolutionBoard, 3, 2)
}

func Example2() {
	var board [][]int = [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 13, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 15, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16},
	}
	solver4 := sudokusolver.NewSolver(4)
	solver4.Solve(board)
	fmt.Printf("Sudoku solved in %v\n", solver4.SolvingTime)
	prettyPrint(solver4.SolutionBoard, 4, 3)
}

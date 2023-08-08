package solver

import (
	"fmt"
	"sudoku/sudoku"
)

type gridNumber struct {
	grid [9][9]int
	n    int
}

func Solver(s sudoku.Sudoku, ch chan [9][9]int) {
	fromOne := true
	auxInt := 0
	for {
		change := false
		fmt.Println("Starting another")
		if s.Solved() {
			ch <- s.GetGrid()
			return
		}
		r, c := s.GetEmptyCell()
		fmt.Println("Row: ", r, "  Col: ", c)
		if fromOne {
			for i := 1; i < 10; i++ {
				if !s.Occupied(r, c) && s.CorrectPlace(i, r, c) {
					s.AddNumber(i, r, c)
					change = true
					fmt.Println("Added Number")
					break
				}
			}
		} else {
			for i := auxInt; i < 10; i++ {
				if !s.Occupied(r, c) && s.CorrectPlace(i, r, c) {
					s.AddNumber(i, r, c)
					change = true
					fromOne = true
					fmt.Println("Added Number")
					break
				}
			}
		}
		if !change {
			auxInt = s.Remove() + 1
			fromOne = false
			fmt.Println("fromOne false")
		}
	}
}

package solver

import "sudoku/sudoku"

type gridNumber struct {
	grid [9][9]int
	n    int
}

func Solver(s sudoku.Sudoku, ch chan [9][9]int) {
	fromOne := true
	auxInt := 0
	for {
		change := false
		if s.Solved() {
			ch <- s.GetGrid()
			return
		}
		r, c := s.GetEmptyCell()
		if fromOne {
			for i := 1; i < 10; i++ {
				if !s.Occupied(r, c) && s.CorrectPlace(i, r, c) {
					s.AddNumber(i, r, c)
					change = true
					break
				}
			}
		} else {
			for i := auxInt; i < 10; i++ {
				if !s.Occupied(r, c) && s.CorrectPlace(i, r, c) {
					s.AddNumber(i, r, c)
					change = true
					fromOne = true
					break
				}
			}
		}
		if !change {
			auxInt = s.Remove() + 1
			fromOne = false
		}
	}
}

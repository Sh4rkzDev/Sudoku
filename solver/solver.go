package solver

import (
	"sudoku/sudoku"
)

func Solver(s sudoku.Sudoku, ch chan [9][9]int) {
	if s.Solved() {
		ch <- s.GetGrid()
		return
	}
	r, c := s.GetEmptyCell()
	for i := 1; i < 10; i++ {
		if !s.Occupied(r, c) && s.CorrectPlace(i, r, c) {
			s.AddNumber(i, r, c)
			break
		}
	}
	Solver(s, ch)
}

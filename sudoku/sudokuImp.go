package sudoku

import (
	stack "github.com/golang-collections/collections/stack"
)

type sudoku struct {
	rows map[int]map[int]bool
	cols map[int]map[int]bool
	grid [9][9]int
	last *stack.Stack
}

func CreateSudoku() Sudoku {
	rows := map[int]map[int]bool{}
	cols := map[int]map[int]bool{}
	for i := 1; i < 10; i++ {
		rows[i] = map[int]bool{}
		cols[i] = map[int]bool{}
	}
	grid := [9][9]int{}
	stk := stack.New()
	return &sudoku{rows, cols, grid, stk}
}

func (s *sudoku) AddNumber(number, row, col int) error {
	if s.Occupied(row, col) {
		panic("The given cell is already occupied")
	}
	if !s.CorrectPlace(number, row, col) {
		panic("The number is already on the given row or column")
	}
	s.rows[row][number] = true
	s.cols[col][number] = true
	s.grid[row-1][col-1] = number
	aux := [3]int{number, row, col}
	s.last.Push(aux)
	return nil
}

func (s *sudoku) Remove() int {
	if s.last.Len() == 0 {
		panic("No number added")
	}
	l := s.last.Pop().([3]int)
	number, row, col := l[0], l[1], l[2]
	delete(s.cols[col], number)
	delete(s.rows[row], number)
	s.grid[row-1][col-1] = 0
	return number
}

func (s *sudoku) Occupied(row, col int) bool {
	return s.grid[row-1][col-1] != 0
}

func (s *sudoku) CorrectPlace(number, row, col int) bool {
	_, aux1 := s.rows[row][number]
	if aux1 {
		return false
	}
	_, aux2 := s.cols[col][number]
	if aux2 {
		return false
	}
	rowFamily, colFamily := (row-1)/3, (col-1)/3
	return checkFam(s.grid, number, rowFamily, colFamily)
}

func checkFam(grid [9][9]int, n, r, c int) bool {
	for i := r * 3; i < (r*3)+3; i++ {
		for j := c * 3; j < (c*3)+3; j++ {
			if grid[i][j] == n {
				return false
			}
		}
	}
	return true
}

func (s *sudoku) GetEmptyCell() (int, int) {
	if s.Solved() {
		panic("The sudoku is already complete")
	}
	for r, row := range s.grid {
		for c, col := range row {
			if col == 0 {
				return r + 1, c + 1
			}
		}
	}
	return -1, -1
}

func (s *sudoku) GetGrid() [9][9]int {
	return s.grid
}

func (s *sudoku) Solved() bool {
	for _, row := range s.grid {
		for _, col := range row {
			if col == 0 {
				return false
			}
		}
	}
	return true
}

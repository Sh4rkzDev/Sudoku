package sudoku

import (
	"fmt"

	stack "github.com/golang-collections/collections/stack"
)

type sudoku struct {
	rows      map[int]map[int]bool
	cols      map[int]map[int]bool
	permStack stack.Stack
	quant     map[int]int
	grid      [9][9]int
	last      *stack.Stack
}

func CreateSudoku() Sudoku {
	rows := map[int]map[int]bool{}
	cols := map[int]map[int]bool{}
	perm := stack.New()
	quant := map[int]int{}
	for i := 1; i < 10; i++ {
		rows[i] = map[int]bool{}
		cols[i] = map[int]bool{}
	}
	grid := [9][9]int{}
	stk := stack.New()
	return &sudoku{rows, cols, *perm, quant, grid, stk}
}

func (s *sudoku) AddNumber(number, row, col int) {
	if s.Occupied(row, col) {
		panic("The given cell is already occupied")
	}
	if !s.CorrectPlace(number, row, col) {
		panic("The number is already on the given row or column")
	}
	s.rows[row][number] = true
	s.cols[col][number] = true
	if s.permStack.Len() != 0 {
		s.permStack.Peek().([]int)[0]++
	}
	s.quant[number]++
	s.grid[row-1][col-1] = number
	aux := [3]int{number, row, col}
	s.last.Push(aux)
	fmt.Println("ADDED: ", number)
}

func (s *sudoku) Remove() int {
	if s.last.Len() == 0 {
		panic("No number added")
	}
	l := s.last.Pop().([3]int)
	number, row, col := l[0], l[1], l[2]
	if s.permStack.Len() != 0 {
		s.permStack.Peek().([]int)[0]--
		if s.permStack.Peek().([]int)[0] == 0 {
			ex := s.permStack.Pop().([]int)
			aux := s.grid[ex[2]][ex[3]]
			delete(s.rows[ex[2]], aux)
			delete(s.cols[ex[3]], aux)
			s.rows[ex[2]][ex[1]] = true
			s.cols[ex[3]][ex[1]] = true
			s.grid[ex[2]-1][ex[3]-1] = ex[1]
			s.quant[aux]--
			s.quant[ex[1]]++
			fmt.Println("RMPERM: ", number)
		}
	}
	delete(s.cols[col], number)
	delete(s.rows[row], number)
	s.grid[row-1][col-1] = 0
	s.quant[number]--
	fmt.Println("RM: ", number)
	return number
}

func (s *sudoku) AddNumberPerm(number, row, col int) {
	if !s.CorrectPlace(number, row, col) {
		panic("The number is already on the given row or column")
	}
	if s.Occupied(row, col) {
		aux := s.grid[row-1][col-1]
		auxSl := []int{1, aux, row, col}
		s.permStack.Push(auxSl)
		delete(s.rows[row], aux)
		delete(s.cols[col], aux)
		s.quant[aux]--
	}
	s.rows[row][number] = true
	s.cols[col][number] = true
	s.quant[number]++
	s.grid[row-1][col-1] = number
	fmt.Println("ADDED PERM: ", number)
}

func (s *sudoku) OneLeft() (error, int, int, int) {
	for k, v := range s.quant {
		if v == 8 {
			row := 0
			col := 0
			for rows, n := range s.rows {
				_, exist := n[k]
				if !exist {
					row = rows
					break
				}
			}
			for cols, n := range s.cols {
				_, exist := n[k]
				if !exist {
					col = cols
					break
				}
			}
			return nil, k, row, col
		}
	}
	return ErrorOneLeft{}, -1, -1, -1
}

type ErrorOneLeft struct{}

func (e ErrorOneLeft) Error() string {
	return "There is no number that is one position away of being complete"
}

func (s *sudoku) Occupied(row, col int) bool {
	return s.grid[row-1][col-1] != 0
}

func (s *sudoku) CorrectPlace(number, row, col int) bool {
	_, aux1 := s.rows[row][number]
	if aux1 {
		fmt.Println("Same row")
		return false
	}
	_, aux2 := s.cols[col][number]
	if aux2 {
		fmt.Println("Same col")
		return false
	}
	rowFamily, colFamily := (row-1)/3, (col-1)/3
	return checkFam(s.grid, number, rowFamily, colFamily)
}

func checkFam(grid [9][9]int, n, r, c int) bool {
	for i := r * 3; i < (r*3)+3; i++ {
		for j := c * 3; j < (c*3)+3; j++ {
			if grid[i][j] == n {
				fmt.Println("Same fam")
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

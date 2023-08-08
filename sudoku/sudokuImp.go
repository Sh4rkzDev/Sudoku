package sudoku

type sudoku struct {
	rows map[int]map[int]bool
	cols map[int]map[int]bool
	grid [9][9]int
}

func CreateSudoku() Sudoku {
	rows := map[int]map[int]bool{}
	cols := map[int]map[int]bool{}
	for i := 1; i < 10; i++ {
		rows[i] = map[int]bool{}
		cols[i] = map[int]bool{}
	}
	grid := [9][9]int{}
	return &sudoku{rows, cols, grid}
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
	return nil
}

func (s *sudoku) Occupied(row, col int) bool {
	return s.grid[row-1][col-1] == 0
}

func (s *sudoku) CorrectPlace(number, row, col int) bool {
	_, aux1 := s.rows[row][number]
	_, aux2 := s.cols[col][number]
	return !aux1 && !aux2
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

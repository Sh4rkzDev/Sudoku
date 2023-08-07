package sudoku

type sudoku struct {
	rows map[int]map[int]bool
	cols map[int]map[int]bool
	grid [9][9]int
}

type ErrorPlacement struct{}

func (e ErrorPlacement) Error() string {
	return "The number is already on the given row or column"
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
		return ErrorPlacement{}
	}
	s.rows[row][number] = true
	s.cols[col][number] = true
	s.grid[row][col] = number
	return nil
}

func (s *sudoku) Occupied(row, col int) bool {
	return s.grid[row][col] == 0
}

func (s *sudoku) CorrectPlace(number, row, col int) bool {
	_, aux1 := s.rows[row][number]
	_, aux2 := s.cols[col][number]
	return !aux1 && !aux2
}

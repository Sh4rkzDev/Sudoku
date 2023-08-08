package sudoku

type Sudoku interface {

	//It will add the given number to the specific cell.
	//In case of not being a correct place to put the number or the cell is already Occupied, it will panic.
	AddNumber(number, row, column int) error

	//Returns true in case of being occupied.
	Occupied(row, column int) bool

	//Returns true if it is correct to place the given number at that cell.
	//It is correct to place a number when the same number is not at the given row and the give column.
	CorrectPlace(number, row, column int) bool

	//Returns the row and column of the first empty cell found. In case of being completed, it will panic.
	GetEmptyCell() (int, int)

	//Returns the grid of the sudoku at that moment. Zeros are equivalents to an empty cell.
	GetGrid() [9][9]int

	//Returns true if the solution is found
	Solved() bool
}

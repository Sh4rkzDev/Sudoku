package sudoku

type Sudoku interface {

	//It will add the given number to the specific cell.
	//In case of not being a correct place to put the number or the cell is already Occupied, it will panic.
	AddNumber(number, row, column int)

	//It will remove the last number placed and return it.
	//In case of no added number, it will panic.
	Remove() int

	//It will add the given number to the specific cell temporary-permanently.
	//This operation will automatically be undo when the operation of putting the original number there is undo (in case of already being occupied).
	//In case of not being occupied originally, it will be undo when undoing for first time an operation that involves the same number.
	//In case of being occupied, it will overwrite the number that is set.
	//In case of not being a correct place to put the number, it will panic.
	AddNumberPerm(number, row, column int)

	//It will return the number that only one position is left to put it, and the row and column where it should be.
	//In case of no number is one position left to be completed, it will return an error too.
	OneLeft() (error, int, int, int)

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

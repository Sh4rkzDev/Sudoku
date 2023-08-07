package sudoku

type Sudoku interface {

	//It will add the given number to the specific cell. In case of being occupied, it will panic.
	//In case of not being a correct place to put the number, it will return an error.
	AddNumber(number, row, column int) error

	//Returns true in case of being occupied.
	Occupied(row, column int) bool

	//Returns true if it is correct to place the given number at that cell.
	//It is correct to place a number when the same number is not at the given row and the give column.
	CorrectPlace(number, row, column int) bool
}

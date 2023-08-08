package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"sudoku/sudoku"
)

func createReq(n int) string {
	switch n {
	case 1:
		return "1st line: "
	case 2:
		return "2nd line: "
	case 3:
		return "3rd line: "
	case 4:
		return "4th line: "
	case 5:
		return "5th line: "
	case 6:
		return "6th line: "
	case 7:
		return "7th line: "
	case 8:
		return "8th line: "
	default:
		return "9th line: "
	}
}

func ReadInput() sudoku.Sudoku {
	input := bufio.NewScanner(os.Stdin)
	s := sudoku.CreateSudoku()
	fmt.Println("Please, enter the numbers with the following format (row line):")
	fmt.Println("1,2,3,,,7,,5,")
	fmt.Println("No number represents blank cell")
	fmt.Println()
	for row := 1; row < 10; row++ {
		fmt.Println(createReq(row))
		input.Scan()
		numsS := input.Text()
		nums := strings.Split(numsS, ",")
		for column, n := range nums {
			if n == "" {
				continue
			}
			nInt, err := strconv.Atoi(n)
			if err != nil {
				panic("Enter a valid number")
			}
			s.AddNumber(nInt, row, column+1)
		}
		fmt.Println()
	}
	return s
}

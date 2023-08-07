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
	for i := 0; i < 9; i++ {
		fmt.Println(createReq(i + 1))
		input.Scan()
		numsS := input.Text()
		nums := strings.Split(numsS, ",")
		for _, n := range nums {

		}
	}
}

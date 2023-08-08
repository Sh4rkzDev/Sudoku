package main

import (
	"fmt"
	"sudoku/inputs"
	"sudoku/solver"
)

func main() {
	sudoku := inputs.ReadInput()
	ch := make(chan [9][9]int)
	go solver.Solver(sudoku, ch)
	sol := <-ch
	close(ch)
	fmt.Printf("  -----------------------------------\n")
	for r, row := range sol {
		for c, col := range row {
			if c == 0 {
				fmt.Printf(" | ")
			}
			fmt.Printf(" %d ", col)
			if (c+1)%3 == 0 {
				fmt.Printf(" | ")
			}
		}
		fmt.Printf("\n")
		if (r+1)%3 == 0 {
			fmt.Printf("  -----------------------------------\n")
		}
	}
}

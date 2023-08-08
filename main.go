package main

import (
	"fmt"
	"time"

	"sudoku/inputs"
	"sudoku/solver"
)

func main() {
	sudoku := inputs.ReadInput()
	ch := make(chan [9][9]int)
	start := time.Now()
	go solver.Solver(sudoku, ch)
	sol := <-ch
	end := time.Now()
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
	fmt.Println()
	fmt.Println("It took ", end.Sub(start))
}

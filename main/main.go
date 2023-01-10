package main

import (
	"fmt"
	"os"
	"sudoku"
)

func main() {
	arg := os.Args
	err, errType := sudoku.IsParameterCorrect(arg)
	if err {
		table := sudoku.Creat(arg)
		tableResolu := sudoku.CallSolveSudoku(table)
		if table == tableResolu {
			fmt.Println("sudoku impossible")
		} else {
			sudoku.DisplaySudoku(tableResolu)
		}
	} else {
		fmt.Println(errType)
	}
}

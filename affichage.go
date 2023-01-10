package sudoku

import "fmt"

func DisplaySudoku(tab [9][9]int) {
	fmt.Print("|-------+-------+-------| \n")
	for y, line := range tab {
		for x := 0; x < len(line); x++ {
			if x%3 == 0 {
				fmt.Print("| ")
				fmt.Print(line[x], " ")
			} else {
				fmt.Print(line[x], " ")
			}
		}
		if (y+1)%3 == 0 {
			fmt.Print("| ", "\n")
			fmt.Print("|-------+-------+-------| ")
		} else {
			fmt.Print("| ")
		}
		fmt.Print("\n")
	}
}

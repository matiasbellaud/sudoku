package sudoku

//// verify if there is 2 time the same number in 3*3 square
func EmptySquare(verifNumber int, table [9][9]int, column int, line int) bool {
	stockColumn := column - (column % 3)
	stockLine := line - (line % 3)
	for column = stockColumn; column < stockColumn+3; column++ {
		for line = stockLine; line < stockLine+3; line++ {
			if table[column][line] == verifNumber {
				return false
			}
		}
	}
	return true
}

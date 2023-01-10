package sudoku

// verify if there are 2 time the same number in line
func EmptyLine(verifNumber int, table [9][9]int, column int, line int) bool {
	for indexLine := 0; indexLine < 9; indexLine++ {
		if table[column][indexLine] == verifNumber {
			return false
		}
	}
	return true
}

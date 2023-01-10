package sudoku

// verify if there are 2 time the same number in column
func EmptyColumn(verifNumber int, table [9][9]int, column int, line int) bool {
	for indexColumn := 0; indexColumn < 9; indexColumn++ {
		if table[indexColumn][line] == verifNumber {
			return false
		}
	}
	return true
}

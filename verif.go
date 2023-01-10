package sudoku

func VerifAll(verifNumber int, table [9][9]int, column int, line int) bool {
	tempo := false
	if EmptyColumn(verifNumber, table, column, line) {
		tempo = true
	} else {
		return false
	}
	if EmptyLine(verifNumber, table, column, line) {
		tempo = true
	} else {
		return false
	}
	if EmptySquare(verifNumber, table, column, line) {
		tempo = true
	} else {
		return false
	}
	return tempo
}

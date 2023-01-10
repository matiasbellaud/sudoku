package sudoku

type posibiltyCase struct {
	posibility int
	x          int
	y          int
}

// calculates the number of posibilities to put a number
func NumberOfPosibility(table [9][9]int, column int, line int) int {
	nbPosibilty := 0
	for testNum := 1; testNum <= 9; testNum++ {
		if VerifAll(testNum, table, column, line) {
			nbPosibilty++
		}
	}
	return nbPosibilty
}

// create a list of empty cells with the number of posibiliter(posibility), row(X), column(ygit pu)
func creatListPosibility(tab [9][9]int) []posibiltyCase {
	var nbPosibilty int
	var listOrder []posibiltyCase = []posibiltyCase{}
	for x := 0; x <= 8; x++ {
		for y := 0; y <= 8; y++ {
			if tab[x][y] == 0 {
				nbPosibilty = NumberOfPosibility(tab, x, y)
				listOrder = append(listOrder, posibiltyCase{nbPosibilty, x, y})
			}
		}
	}
	return listOrder
}

func TriListOrder(listOrder []posibiltyCase) []posibiltyCase {
	var listOrderOf2 []posibiltyCase = []posibiltyCase{}
	for _, cell := range listOrder {
		if cell.posibility == 2 {
			listOrderOf2 = append(listOrderOf2, cell)
		}
	}
	return listOrderOf2
}

// Call SolveSudoku
func CallSolveSudoku(tab [9][9]int) [9][9]int {
	listOrder := creatListPosibility(tab)
	tab, _ = SolveSudoku(tab, listOrder, 0)

	return tab

}

// complete the sudoku so that it is good
func SolveSudoku(tab [9][9]int, listOrder []posibiltyCase, index int) ([9][9]int, bool) {
	var tabModif [9][9]int = tab
	var test bool
	if len(listOrder)-1 >= index {
		test = false
		for i := 1; i <= 9; i++ {
			if VerifAll(i, tab, int(listOrder[index].x), int(listOrder[index].y)) {
				tabModif[listOrder[index].x][listOrder[index].y] = i
				test = true
				tabModif, test = SolveSudoku(tabModif, listOrder, index+1)
				if test {
					return tabModif, test
				}
			}
		}
	} else {
		return tabModif, true
	}
	return tab, test
}

package sudoku

func Creat(tab []string) [9][9]int {
	var tabSudoku [9][9]int
	var tempo []rune
	for i := 0; i <= 8; i++ {
		tempo = []rune(tab[i+1])
		for index, num := range tempo {
			if num == '.' {
				tabSudoku[i][index] = 0
			} else {
				tabSudoku[i][index] = int(num - 48)
			}
		}

	}
	return tabSudoku
}

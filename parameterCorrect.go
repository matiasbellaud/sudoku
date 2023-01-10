package sudoku

func IsParameterCorrect(parameterSudo []string) (bool, string) {
	var answer bool = false
	if len(parameterSudo) == 10 {
		for i := 1; i < len(parameterSudo); i++ { //Verify if the lenght of parameter is good and also if the character isn't a space
			if len(parameterSudo[i]) == 9 {
				if parameterSudo[i] == " " {
					return false, "There is a space !"
				}
				for index, letter := range parameterSudo[i] { //If the letter we verify is a number
					if letter >= '1' && letter <= '9' || letter == '.' {
						answer = true
					} else {
						return false, "One character is not a number !"
					}
					for verifIndex, verifLetter := range parameterSudo[i] { //For don't have the two same number in the same string
						if letter == verifLetter && index != verifIndex && letter != '.' {
							return false, "One number is two times in the same string !"
						}
					}

				}
			} else {
				return false, "There are less or more than 9 characters !"
			}
		}
	}
	if answer {
		return answer, "All is good, the parameters are correct !"
	} else {
		return answer, "error"
	}
}

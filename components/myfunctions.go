package components

import "fmt"

const (
	lettersNum         = 26		// number of letters in alphabet
	maxLettersOneRow   = 702	// max number of columns per pow
	unicodeLetterStart = 65		// starting position of the letter 'A' in the unicode table
)

// GetSpreadsheetNotation return a string representation of a cell number is spreadsheet notation by map of alphabet
func GetSpreadsheetNotation(n uint64) (result string) {
	var (
		letters map[uint64]string
		rowNum  uint64
		colNum  uint64
	)

	letters = map[uint64]string{
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "M",
		14: "N",
		15: "O",
		16: "P",
		17: "Q",
		18: "R",
		19: "S",
		20: "T",
		21: "U",
		22: "V",
		23: "W",
		24: "X",
		25: "Y",
		26: "Z",
	}

	rowNum = n / (maxLettersOneRow)

	// If the remainder of dividing by the max number of columns in th row is zero,
	// then we will reduce the position of the raw
	if n % maxLettersOneRow == 0 {
		rowNum -= 1
	}

	result += fmt.Sprintf("%d", rowNum+1)
	colNum = n - (rowNum * maxLettersOneRow)

	if colNum <= lettersNum {
		result += letters[colNum]
	} else {
		firstLetterNum := (colNum - 1) / lettersNum
		lastLetterNum := colNum - (firstLetterNum*lettersNum)
		result += letters[firstLetterNum] + letters[lastLetterNum]
	}

	return result
}

// GetSpreadsheetNotation return a string representation of a cell number is spreadsheet notation by unicode table
func GetSpreadsheetNotation2(n uint64) (result string) {
	var (
		rowNum uint64
		colNum uint64
	)

	rowNum = n / maxLettersOneRow

	// If the remainder of dividing by the max number of columns in th row is zero,
	// then we will reduce the position of the raw
	if n % maxLettersOneRow == 0 {
		rowNum -= 1
	}

	result += fmt.Sprintf("%d", rowNum+1)
	colNum = n - (rowNum * maxLettersOneRow)

	if colNum <= lettersNum {
		result += string(unicodeLetterStart + colNum - 1)
	} else {
		firstLetterNum := ((colNum - 1) / lettersNum) - 1
		lastLetterNum := (colNum - ((firstLetterNum + 1) * lettersNum)) - 1
		result += string(unicodeLetterStart+firstLetterNum) + string(unicodeLetterStart+lastLetterNum)
	}

	return result
}


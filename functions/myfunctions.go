package functions

import (
	"../constants"
	"fmt"
)

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

	rowNum = n / constants.MaxLettersOneRow
	result += fmt.Sprintf("%d", rowNum+1)
	colNum = n - (rowNum * constants.MaxLettersOneRow)

	if colNum <= constants.LettersNum {
		result += letters[colNum]
	} else {
		firstLetterNum := colNum / constants.LettersNum
		lastLetterNum := colNum - firstLetterNum * constants.LettersNum
		result += letters[firstLetterNum] + letters[lastLetterNum]
	}

	return result
}

func GetSpreadsheetNotation2(n uint64) (result string) {
	var (
		rowNum uint64
		colNum uint64
	)

	rowNum = n / constants.MaxLettersOneRow
	result += fmt.Sprintf("%d", rowNum+1)
	colNum = n - (rowNum * constants.MaxLettersOneRow)

	if colNum <= constants.LettersNum {
		result += string(constants.UnicodeLetterStart + colNum - 1)
	} else {
		firstLetterNum := (colNum / constants.LettersNum) - 1
		lastLetterNum := colNum - ((firstLetterNum + 1) * constants.LettersNum)
		result += string(constants.UnicodeLetterStart + firstLetterNum) + string(constants.UnicodeLetterStart + lastLetterNum)
	}

	return result
}

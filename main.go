package main

// ДИМАСЯ НЕ КРИЧИ! Я ПОТОМ РАСКИДАЮ ПО ФАЙЛАМ

import (
	"fmt"
	"time"
)

const (
	lettersNum         = 26
	maxLettersOneRow   = 702
	unicodeLetterStart = 65
)

func getSpreadsheetNotation(n uint64) (result string) {
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

	rowNum = n / maxLettersOneRow
	result += fmt.Sprintf("%d", rowNum+1)
	colNum = n - (rowNum * maxLettersOneRow)

	if colNum <= lettersNum {
		result += letters[colNum]
	} else {
		firstLetterNum := colNum / lettersNum
		lastLetterNum := colNum - firstLetterNum*lettersNum
		result += letters[firstLetterNum] + letters[lastLetterNum]
	}

	return result
}

func getSpreadsheetNotation2(n uint64) (result string) {
	var (
		rowNum uint64
		colNum uint64
	)

	rowNum = n / maxLettersOneRow
	result += fmt.Sprintf("%d", rowNum+1)
	colNum = n - (rowNum * maxLettersOneRow)

	if colNum <= lettersNum {
		result += string(unicodeLetterStart + colNum - 1)
	} else {
		firstLetterNum := (colNum / lettersNum) - 1
		lastLetterNum := colNum - ((firstLetterNum + 1) * lettersNum)
		result += string(unicodeLetterStart+firstLetterNum) + string(unicodeLetterStart+lastLetterNum)
	}

	return result
}

func main() {
	start := time.Now()
	for i := 1; i < 100000; i++ {
		getSpreadsheetNotation(uint64(i))
	}
	//fmt.Println(getSpreadsheetNotation(27))
	//fmt.Println(getSpreadsheetNotation(1405))
	elapsed := time.Since(start)
	fmt.Printf("getSpreadsheetNotation took %s\n", elapsed)

	start = time.Now()
	for i := 1; i < 100000; i++ {
		getSpreadsheetNotation2(uint64(i))
	}
	// fmt.Println(getSpreadsheetNotation2(27))
	// fmt.Println(getSpreadsheetNotation2(1405))
	elapsed = time.Since(start)
	fmt.Printf("getSpreadsheetNotation2 took %s\n", elapsed)

	for i := 1; i < 68; i++ {
		fmt.Println(getSpreadsheetNotation2(uint64(i)), " ->", i)
	}
}

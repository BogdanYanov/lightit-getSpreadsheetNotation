package main

import (
	"fmt"
	"github.com/BogdanYanov/lightit-getSpreadsheetNotation/components"
)

func main() {
	var inputNum uint64 = 702
	fmt.Printf("Input num -> %d; Result -> %s\n", inputNum, components.GetSpreadsheetNotation(inputNum))
	fmt.Printf("Input num -> %d; Result -> %s\n", inputNum, components.GetSpreadsheetNotation2(inputNum))
	fmt.Printf("Input num -> %d; Result -> %s\n", inputNum, components.GetSpreadsheetNotationFor(inputNum))
}

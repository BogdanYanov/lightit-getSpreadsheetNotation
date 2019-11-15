package main

// ДИМАСЯ НЕ КРИЧИ! Я ПОТОМ РАСКИДАЮ ПО ФАЙЛАМ

import (
	"../functions"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i < 100000; i++ {
		functions.GetSpreadsheetNotation(uint64(i))
	}
	//fmt.Println(getSpreadsheetNotation(27))
	//fmt.Println(getSpreadsheetNotation(1405))
	elapsed := time.Since(start)
	fmt.Printf("getSpreadsheetNotation took %s\n", elapsed)

	start = time.Now()
	for i := 1; i < 100000; i++ {
		functions.GetSpreadsheetNotation2(uint64(i))
	}
	// fmt.Println(getSpreadsheetNotation2(27))
	// fmt.Println(getSpreadsheetNotation2(1405))
	elapsed = time.Since(start)
	fmt.Printf("getSpreadsheetNotation2 took %s\n", elapsed)

	for i := 1; i < 68; i++ {
		fmt.Println(functions.GetSpreadsheetNotation(uint64(i)), " ->", i)
	}

	for i := 1; i < 68; i++ {
		fmt.Println(functions.GetSpreadsheetNotation2(uint64(i)), " ->", i)
	}
}

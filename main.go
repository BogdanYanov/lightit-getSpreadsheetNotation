package main

import (
	"fmt"
	"github.com/BogdanYanov/lightit-getSpreadsheetNotation/components"
	"time"
)

func main() {
	var (
		start   time.Time
		expired time.Duration
	)

	start = time.Now()

	for i := 1; i < 10000; i++ {
		components.GetSpreadsheetNotation(uint64(i))
	}

	expired = time.Since(start)

	fmt.Printf("GetSpreadsheetNotation took %s\n", expired)

	start = time.Now()

	for i := 1; i < 10000; i++ {
		components.GetSpreadsheetNotation2(uint64(i))
	}

	expired = time.Since(start)

	fmt.Printf("GetSpreadsheetNotation2 took %s\n", expired)
}
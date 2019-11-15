# Spreadsheet Notation Conversion

A spreadsheet is an application that allows for the organization of data in rows and columns of a table. Each rectangle within a table is called a cell. We define the following spreadsheet conversitions:
- Rows are labeled sequentially from top to bottom, starting with row 1 and ending with row 10^9
- Columns are labeled sequentially from left to right, start with column A and ending column ZZ, where Z is followed by AA, AA followed by AB, and so on. As a result there are exactly 26x27 = 702 columns per row. Its important to notate that column labels are always capitalized.
- A cell located at the intersection of row R and column C is written as RC in spreadsheet notation. For example, the cell at the intersection of row 7 and column AH is written as 7AH

We define a cell's cell number to be the long integer assigned to it when the cells are counted sequentially from left to right and top to bottom.

**Function Description**

Complete the function getSpreadsheetNotation. The function must return a string representation of a cell number is spreadsheet notation.

```
func getSpreadsheetNotation(n uint64) (result string) {}
```

**Constraints**

1 <= n <= 10^12

**Sample Case 0**

Input: 27
Output: 1AA

**Sample Case 1**

Input: 1405
Output: 3A

# How run the programm

Type in project directory:
```
go run main.go
```
On display you will be see the next message:
```
GetSpreadsheetNotation took 58.502796ms
GetSpreadsheetNotation2 took 5.168334ms
```
This is the time during which these two functions calculate the cell by the number from 1 to 10000

## Tests

If you go to the folder **components/** you can start the test. 

**Test pairs**

- 1 must return "1A"
- 26 must return "1Z"
- 27 must return "1AA"
- 52 must return "1AZ"
- 53 must return "1BA"
- 702 must return "1ZZ"
- 1405 must return "3A"
- 1430 must return "3Z"
- 1431 must return "3AA"

```
~.../components$ go test
PASS
ok      github.com/BogdanYanov/lightit-getSpreadsheetNotation/components        0.004s
```

```
~.../components$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/BogdanYanov/lightit-getSpreadsheetNotation/components
BenchmarkGetSpreadsheetNotation-2         290673              4064 ns/op
BenchmarkGetSpreadsheetNotation2-2       2340072               505 ns/op
PASS
ok      github.com/BogdanYanov/lightit-getSpreadsheetNotation/components        3.881s
```

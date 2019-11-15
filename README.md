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

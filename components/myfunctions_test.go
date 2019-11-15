package components

import "testing"

type testPair struct {
	value  uint64
	result string
}

var tests = []testPair{
	{1, "1A"},
	{26, "1Z"},
	{27, "1AA"},
	{52, "1AZ"},
	{53, "1BA"},
	{1405, "3A"},
	{1430, "3Z"},
	{1431, "3AA"},
}

func TestGetSpreadsheetNotation(t *testing.T) {
	for _, pair := range tests {
		v := GetSpreadsheetNotation(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestGetSpreadsheetNotation2(t *testing.T) {
	for _, pair := range tests {
		v := GetSpreadsheetNotation2(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

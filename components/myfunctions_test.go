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
	{702, "1ZZ"},
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

func BenchmarkGetSpreadsheetNotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetSpreadsheetNotation(702)
	}
}

func BenchmarkGetSpreadsheetNotation2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetSpreadsheetNotation2(702)
	}
}
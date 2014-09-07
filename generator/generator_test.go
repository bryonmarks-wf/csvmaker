package generator

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	checkResult(t, "0\n", 1, 1)
	checkResult(t, "0, 0\n1, 1\n", 2, 2)
	checkResult(t, "0, 0, 0\n1, 1, 1\n", 2, 3)
}

func checkResult(t *testing.T, expected string, rows int, cols int) {
	csvReader := New(rows, cols)
	result := csvReader.buffer.String()
	if csvReader.buffer.String() != expected {
		t.Errorf(`Expected Read() to return %s but got %s`, expected, result)
		t.Fail()
	}
}


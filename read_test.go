package md

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	filename := "tests/mock_data/mock1.md"

	fileLines, err := ParseFile(filename)

	if err != nil {
		t.Log("Invalid file")
		t.Fail()
	}

	if len(fileLines) < 1 {
		t.Log("File should contain data")
		t.Fail()
	}
}

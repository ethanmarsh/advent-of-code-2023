package filereader

import (
	"testing"
)

func TestFileReader(t *testing.T) {
	lines := ReadFile("test_file.txt")
	if lines[0][0] != "5" {
		t.Error("Failed to read first line of text from file")
	}
	if lines[1][1] != "4" {
		t.Error("Failed to read second line of text from file")
	}
	if lines[2][2] != "3" {
		t.Error("Failed to read third line of text from file")
	}
}

package filereader

import (
	"testing"
)

func TestFileReader(t *testing.T) {
	lines := ReadFile("test_file.txt")
	if lines[0] != "Game 1" {
		t.Error("Failed to read first line of text from file")
	}
	if lines[1] != "Game 2" {
		t.Error("Failed to read second line of text from file")
	}
	if lines[2] != "Game 3" {
		t.Error("Failed to read third line of text from file")
	}
}

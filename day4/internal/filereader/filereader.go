package filereader

import (
	"bufio"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) [][]string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var splitLines [][]string
	for _, line := range lines {
		splitLine := strings.Split(line, "")
		splitLines = append(splitLines, splitLine)
	}

	return splitLines
}

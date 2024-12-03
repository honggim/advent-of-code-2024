package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name          string
		filepath      string
		expectedPart1 int
		expectedPart2 int
	}{
		{
			name:          "example",
			filepath:      "./example.txt",
			expectedPart1: 11,
			expectedPart2: 31,
		},
		{
			name:          "input",
			filepath:      "./input.txt",
			expectedPart1: 1579939,
			expectedPart2: 20351745,
		},
	}

	const separator = "   "
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left, right, err := parseFile(tt.filepath, separator)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// part 1
			actual, err := part1(left, right)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if actual != tt.expectedPart1 {
				t.Errorf("expected: %d, actual: %d", tt.expectedPart1, actual)
			}

			// part 2
			actual = part2(left, right)
			if actual != tt.expectedPart2 {
				t.Errorf("expected: %d, actual: %d", tt.expectedPart2, actual)
			}
		})
	}
}

// TODO: move into diff directory
func parseFile(filepath, separator string) ([]int, []int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}

	left := make([]int, 0)
	right := make([]int, 0)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		l, r, err := parseLine(line, separator)
		if err != nil {
			return nil, nil, err
		}
		left = append(left, l)
		right = append(right, r)
	}

	return left, right, nil
}

func parseLine(line, separator string) (int, int, error) {
	const expectedLen = 2
	words := strings.Split(line, separator)
	if len(words) != expectedLen {
		return 0, 0, fmt.Errorf("len(words), expected %d, actual: %d", expectedLen, len(words))
	}

	left, _ := strconv.Atoi(words[0])
	right, _ := strconv.Atoi(words[1])

	return left, right, nil
}

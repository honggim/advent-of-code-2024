package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	const separator = " "
	tests := []struct {
		name     string
		filepath string
		expected int
	}{
		{
			name:     "example",
			filepath: "./example.txt",
			expected: 2,
		},
		{
			name:     "input",
			filepath: "./input.txt",
			expected: 598,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := parseFile(tt.filepath, separator)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			actual := part1(lines)
			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}

func parseFile(filepath, separator string) ([][]int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	lines := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		nums, err := parseLine(line, separator)
		if err != nil {
			return nil, err
		}
		lines = append(lines, nums)
	}

	return lines, nil
}

func parseLine(line, separator string) ([]int, error) {
	words := strings.Split(line, separator)

	nums := make([]int, len(words))
	for i, word := range words {
		num, _ := strconv.Atoi(word)
		nums[i] = num
	}

	return nums, nil
}

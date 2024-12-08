package day07

import (
	"advent-of-code-2024/pkg/parser"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestDay07(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		// fn       func([]string) int
		fn       func(int, []int) int
		expected int
	}{
		{
			name:     "part 1: example.txt",
			filepath: "./example.txt",
			fn:       part1,
			expected: 3749,
		},
		{
			name:     "part 1: input.txt",
			filepath: "./input.txt",
			fn:       part1,
			expected: 8401132154762,
		},
		{
			name:     "part 2: example.txt",
			filepath: "./example.txt",
			fn:       part2,
			expected: 11387,
		},
		{
			name:     "part 2: input.txt",
			filepath: "./input.txt",
			fn:       part2,
			expected: 95297119227552,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := parser.Parse(tt.filepath)
			if err != nil {
				t.Errorf("unexpected err: %v", err)
			}

			actual := 0
			for _, line := range lines {
				sum, sentence := parseLine(line)
				nums := convertToInts(sentence, " ")

				actual += tt.fn(sum, nums)
			}

			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}

// assume well formed input
// line: "83: 15 6"
func parseLine(line string) (int, string) {
	tokens := strings.Split(line, ": ")

	sum, _ := strconv.Atoi(tokens[0])

	return sum, tokens[1]
}

func TestHelpers(t *testing.T) {
	sum, sentence := parseLine("83: 15 6")
	if sum != 83 {
		t.Errorf("expected: %d, actual: %d", 83, sum)
	}

	nums := convertToInts(sentence, " ")
	if nums[0] != 15 || nums[1] != 6 {
		t.Errorf("error with convertToInts")
	}

	result := int(math.Pow(2, 3))
	if result != 8 {
		t.Errorf("exponential syntax: %d", result)
	}
}

package day02

import (
	"advent-of-code-2024/pkg/math"
)

func part1(levels [][]int) int {
	sum := 0
	for _, level := range levels {
		if evaluate(level) {
			sum++
		}
	}
	return sum
}

func evaluate(level []int) bool {
	var xor bool
	for i := 0; i < len(level)-1; i++ {
		if i == 0 {
			xor = !math.IsNegative(level[i], level[i+1])
		}
		if xor != !math.IsNegative(level[i], level[i+1]) {
			return false
		}

		abs := math.AbsDiff(level[i], level[i+1])
		if abs == 0 || abs > 3 {
			return false
		}
	}

	return true
}

func part2(levels [][]int) int {
	sum := 0
	for _, level := range levels {
		if evaluate(level) {
			sum++
			continue
		}

		newLevels := permutations(level)
		for _, newLevel := range newLevels {
			if evaluate(newLevel) {
				sum++
				break
			}
		}
	}
	return sum
}

func permutations(level []int) [][]int {
	matrix := make([][]int, len(level))

	for i := range level {
		row := make([]int, 0)

		for j, val := range level {
			if i != j {
				row = append(row, val)
			}
		}

		matrix[i] = row
	}

	return matrix
}

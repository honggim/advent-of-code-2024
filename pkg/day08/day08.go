package day08

import (
	"fmt"
)

type Point struct {
	row, col int
}

func (p Point) Token() string {
	return fmt.Sprintf("%d|%d", p.row, p.col)
}

func part1(puzzle [][]rune) int {
	rowMax := len(puzzle) - 1
	colMax := len(puzzle[0]) - 1
	hash := make(map[string]bool)

	for _, antennaes := range mapAntennaToLocations(puzzle) {
		locations := antinodes(antennaes)
		fmt.Println(locations)
		for _, location := range locations {
			if location.col < 0 || location.row < 0 {
				continue
			}
			if location.col > colMax || location.row > rowMax {
				continue
			}

			hash[location.Token()] = true
		}
	}
	return len(hash)
}

func mapAntennaToLocations(puzzle [][]rune) map[rune][]Point {
	hash := make(map[rune][]Point)

	for row, line := range puzzle {
		for col, char := range line {
			if char == '.' {
				continue
			}

			p := Point{row: row, col: col}
			if points, ok := hash[char]; ok {
				points = append(points, p)
				hash[char] = points
			} else {
				hash[char] = []Point{p}
			}
		}
	}

	return hash
}

func antinodes(antennas []Point) []Point {
	locations := make([]Point, 0)

	for i := 0; i < len(antennas); i++ {
		first := antennas[i]
		for j := i + 1; j < len(antennas); j++ {
			second := antennas[j]

			locations = append(locations, findAntinodeLocations(first, second)...)
		}
	}

	return locations
}

func findAntinodeLocations(a, b Point) []Point {
	colDiff := a.col - b.col
	rowDiff := a.row - b.row

	aPrime := a
	aPrime.col += colDiff
	aPrime.row += rowDiff

	bPrime := b
	bPrime.col -= colDiff
	bPrime.row -= rowDiff

	return []Point{aPrime, bPrime}
}

package day03

import (
	"regexp"
	"strconv"
)

func part1(input string) int {
	sum := 0

	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(input, -1)

	for _, match := range matches {
		sum += mul(match)
	}

	return sum
}

// assume well formed input
func mul(match string) int {
	r, _ := regexp.Compile(`\d+`)
	nums := r.FindAllString(match, -1)

	num0, _ := strconv.Atoi(nums[0])
	num1, _ := strconv.Atoi(nums[1])

	return num0 * num1
}

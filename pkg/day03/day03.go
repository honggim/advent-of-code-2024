package day03

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	reMul, _ = regexp.Compile(`mul\(\d+,\d+\)`)
	reNum, _ = regexp.Compile(`\d+`)
)

func part1(input string) int {
	sum := 0

	matches := reMul.FindAllString(input, -1)

	for _, match := range matches {
		sum += mul(match)
	}

	return sum
}

// assume well formed input
func mul(match string) int {
	nums := reNum.FindAllString(match, -1)

	num0, _ := strconv.Atoi(nums[0])
	num1, _ := strconv.Atoi(nums[1])

	return num0 * num1
}

func part2(input string) int {
	sum := 0

	lines := strings.Split(input, `don't()`)
	for i, line := range lines {
		if i == 0 {
			sum += part1(line)
			continue
		}

		index := strings.Index(line, `do()`)
		if index == -1 {
			continue
		}

		// len(`do()`) == 4
		sum += part1(line[index+4:])
	}

	return sum
}

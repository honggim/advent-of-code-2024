package day07

import (
	"fmt"
	"strconv"
	"strings"
)

// line: 3267: 81 40 27
// either add or multiply between operations, left to right
// if equal, then return. else 0
func part1(lines []string) int {
	total := 0

	for _, line := range lines {
		sum, sentence := parseLine(line)
		nums := convertToInts(sentence, " ")

		total += checkPart1(sum, nums)
	}

	return total
}

func part2(lines []string) int {
	total := 0

	for _, line := range lines {
		sum, sentence := parseLine(line)
		nums := convertToInts(sentence, " ")

		total += checkPart2(sum, nums)
	}

	return total
}

// assume well formed input
// line: "83: 15 6"
func parseLine(line string) (int, string) {
	tokens := strings.Split(line, ": ")
	// assert.Len(tokens, 2)

	sum, _ := strconv.Atoi(tokens[0])

	return sum, tokens[1]
}

// TODO: move outside package
func convertToInts(sentence, separator string) []int {
	words := strings.Split(sentence, separator)
	nums := make([]int, len(words))

	for i, word := range words {
		nums[i], _ = strconv.Atoi(word)
	}

	return nums
}

func checkPart1(sum int, nums []int) int {
	queue := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		//TODO: newQueue := make([]int, len(queue) * 2)
		newQueue := make([]int, 0)

		for _, prev := range queue {
			//TODO: optimize if val > sum
			newQueue = append(newQueue, prev+num)
			newQueue = append(newQueue, prev*num)
		}

		queue = newQueue
	}

	for _, num := range queue {
		if num == sum {
			return sum
		}
	}
	return 0
}

func checkPart2(sum int, nums []int) int {
	queue := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		//TODO: newQueue := make([]int, len(queue) * 2)
		newQueue := make([]int, 0)

		for _, prev := range queue {
			//TODO: optimize if val > sum
			newQueue = append(newQueue, prev+num)
			newQueue = append(newQueue, prev*num)
			newQueue = append(newQueue, concatInts(prev, num))
		}

		queue = newQueue
	}

	for _, num := range queue {
		if num == sum {
			return sum
		}
	}
	return 0
}

func concatInts(left, right int) int {
	val := fmt.Sprintf("%d%d", left, right)
	num, _ := strconv.Atoi(val)

	return num
}

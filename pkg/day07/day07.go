package day07

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(sum int, nums []int) int {
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

func part2(sum int, nums []int) int {
	queue := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		newQueue := make([]int, 0)

		for _, prev := range queue {
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

// TODO: move outside package
func convertToInts(sentence, separator string) []int {
	words := strings.Split(sentence, separator)
	nums := make([]int, len(words))

	for i, word := range words {
		nums[i], _ = strconv.Atoi(word)
	}

	return nums
}

// TODO: move outside package
func concatInts(left, right int) int {
	val := fmt.Sprintf("%d%d", left, right)
	num, _ := strconv.Atoi(val)

	return num
}

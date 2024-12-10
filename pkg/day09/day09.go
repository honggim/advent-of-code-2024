package day09

import (
	"strconv"
	"strings"
)

func part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += calculateChecksum(line)
	}

	return sum
}

func calculateChecksum(line string) int {
	// 2333133121414131402
	nums := convertToIntArray(line)

	// 00...111...2...333.44.5555.6666.777.888899
	disk := convertToDisk(nums)
	// fmt.Println(strings.Join(disk, ""))

	// 0099811188827773336446555566..............
	defragDisk(disk)
	// fmt.Println(strings.Join(disk, ""))

	return checksum(disk)
}

func convertToIntArray(line string) []int {
	tokens := strings.Split(line, "")
	nums := make([]int, len(tokens))

	for i, token := range tokens {
		nums[i], _ = strconv.Atoi(token)
	}

	return nums
}

const emptyData = "."

func convertToDisk(nums []int) []string {
	output := make([]string, 0)

	for index, num := range nums {
		val := emptyData
		if index%2 == 0 {
			val = strconv.Itoa(index / 2)
		}
		output = append(output, getDisk(val, num)...)
	}

	return output
}

func getDisk(index string, num int) []string {
	output := make([]string, num)

	for i := 0; i < num; i++ {
		output[i] = index
	}

	return output
}

func defragDisk(disk []string) {
	first := 0
	last := len(disk) - 1

	for first < last {
		if disk[first] == emptyData && disk[last] != emptyData {
			disk[first], disk[last] = disk[last], disk[first]
		}

		if disk[first] != emptyData {
			first++
		}

		if disk[last] == emptyData {
			last--
		}
	}
}

func checksum(disk []string) int {
	sum := 0

	for i, data := range disk {
		if data == emptyData {
			return sum
		}
		num, _ := strconv.Atoi(data)
		sum += num * i
	}

	return sum
}

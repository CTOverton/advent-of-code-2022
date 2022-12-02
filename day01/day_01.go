package day01

import (
	"advent-of-code-2022/utils"
	"sort"
	"strconv"
)

type elf struct {
	items []int
	total int
}

func Day01() {
	input := utils.ReadInputFile("day01/input.txt")

	var elves []elf
	var current elf
	item := ""
	sum := 0
	prevNewLine := false

	topDawg := 0

	for _, line := range input {
		if string(line) == "\n" {
			if prevNewLine {
				prevNewLine = false
				current.total = sum
				elves = append(elves, current)

				if sum > topDawg {
					topDawg = sum
				}

				sum = 0
				current = elf{}
			} else {
				i, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}

				current.items = append(current.items, i)
				sum += i
				item = ""
				prevNewLine = true
			}
		} else {
			prevNewLine = false
			item = item + string(line)
		}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].total > elves[j].total
	})

	println("part 1: ", topDawg)
	println("part 2: ", elves[0].total+elves[1].total+elves[2].total)
}

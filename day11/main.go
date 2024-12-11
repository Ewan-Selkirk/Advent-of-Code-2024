package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"strconv"
)

func main() {
	input := load_input()

	part_1(input)
	part_2(input)
}

func load_input() string {
	contents, err := os.ReadFile("input")

	if err != nil {
		path, _ := os.Executable()
		contents, err := os.ReadFile(filepath.Dir(path) + "input")

		if err != nil {
			panic(err)
		}

		return string(contents)
	}

	return string(contents)
}

func string_to_int(number string) int {
	value, err := strconv.Atoi(number)

	if err != nil {
		panic(err)
	}

	return value
}

func warp_stones(stones *map[int]int) {
	temp := make(map[int]int)

	for stone, count := range *stones {
		str := strconv.Itoa(stone)
		if stone == 0 {
			temp[1] += count
		} else if len(str) % 2 == 0 {
			mid := len(str) / 2
			temp[string_to_int(str[:mid])] += count
			temp[string_to_int(str[mid:])] += count
		} else {
			temp[stone * 2024] += count
		}
	}

	// fmt.Println(stones)
	*stones = temp
}

func part_1(input string) {
	stones := make(map[int]int)

	for _, s := range strings.Fields(input) {
		stones[string_to_int(s)]++
	}

	for i := 0; i < 25; i++ {
		warp_stones(&stones)
	}

	count := 0
	for _, c := range stones {
		count += c
	}

	fmt.Println("Part 1:", count)
}


func part_2(input string) {
	stones := make(map[int]int)

	for _, s := range strings.Fields(input) {
		stones[string_to_int(s)]++
	}

	for i := 0; i < 75; i++ {
		warp_stones(&stones)
	}

	count := 0
	for _, c := range stones {
		count += c
	}

	fmt.Println("Part 2:", count)
}

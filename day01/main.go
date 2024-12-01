package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
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

func create_lists(input string) ([]int, []int) {
	fields := strings.Fields(input)
	var left_list []int
	var right_list []int

	for i, id := range fields {
		value, _ := strconv.Atoi(id)

		if value == 0 {
			fmt.Println(i, value)
		}

		if (i % 2 == 0) { 
			left_list = append(left_list, value) 
		} else {
			right_list = append(right_list, value)
		}
	}

	return left_list, right_list
}

func part_1(input string) {
	left_list, right_list := create_lists(input)
	result := 0

	sort.Ints(left_list)
	sort.Ints(right_list)
	
	for i, id := range left_list {
		diff := id - right_list[i]
		if diff < 0 {
			diff = -diff
		}

		result += diff
	}

	fmt.Println(result)
}

func part_2(input string) {
	left_list, right_list := create_lists(input)
	result := 0

	for _, l_id := range left_list {
		counter := 0
		for _, r_id := range right_list {
			if l_id == r_id {
				counter++
			}
		}

		result += counter * l_id
	}

	fmt.Println(result)
}

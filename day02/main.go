package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func string_to_int(number string) int {
	value, err := strconv.Atoi(number)

	if err != nil {
		panic(err)
	}

	return value
}

func part_1(input string) {
	safe := 0
	unsafe := 0

	var changes [][]int

	// Get the changes in values (E.G. i - i + 1) for each value of each lines
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		var numbers []int
		
		for index := 0; index < len(fields) - 1; index++ {
			numbers = append(numbers, string_to_int(fields[index]) - string_to_int(fields[index + 1]))
		}

		changes = append(changes, numbers)
	}

	for _, c := range changes {
		isSafe := true
		isNegative := c[0] < 0
		for _, v := range c {
			if (!isNegative && (v > 0 && v <= 3)) || (isNegative && (v < 0 && v >= -3)) { continue }

			isSafe = false
			unsafe++
			fmt.Println(c, "is unsafe!")
			break
		}

		if isSafe {
			safe++
			fmt.Println(c, "is safe!")
		}
	}

	fmt.Println("Safe Reports:", safe)
	fmt.Println("Unsafe Reports:", unsafe)
}

func part_2(input string) {
	is_valid := func(input []int) bool {
		negative := input[0] < 0

		for _, v := range input {
			if (!negative && (v > 0 && v <= 3)) || (negative && (v < 0 && v >= -3)) { continue }

			return false
		}

		return true
	}

	string_slice_to_int := func(input []string) []int {
		var output []int
		for _, v := range input {
			output = append(output, string_to_int(v))
		}

		return output
	}

	get_diffs := func(input []int) []int {
		var output []int
		for i := 1; i < len(input); i++  {
			output = append(output, input[i - 1] - input[i])
		}

		return output
	}

	safe := 0
	unsafe := 0

	// Get the changes in values (E.G. i - i + 1) for each value of each lines
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		numbers := string_slice_to_int(fields)
		diffs := get_diffs(numbers)

		if is_valid(diffs) {
			safe++
			fmt.Println(line, "is safe!")
			continue
		} else {
			// Soft-locked myself because I assumed the first number couldn't be removed :(
			// I am dumb
			for i := 0; i < len(numbers); i++ {
				changesMinusOne := make([]int, 0, len(numbers) - 1)
				changesMinusOne = append(changesMinusOne, numbers[:i]...)
				changesMinusOne = append(changesMinusOne, numbers[i + 1:]...)

				fmt.Println(get_diffs(changesMinusOne))

				if is_valid(get_diffs(changesMinusOne)) {
					safe++
					fmt.Println(line, "is safe!")
					break
				}

				if i == len(numbers) - 1 {
					unsafe++
					fmt.Println(line, "is unsafe!")
				}
			}
		}
	}

	fmt.Println("Safe Reports:", safe)
	fmt.Println("Unsafe Reports:", unsafe)
}

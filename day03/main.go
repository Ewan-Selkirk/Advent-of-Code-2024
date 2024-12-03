package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"regexp"
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
	counter := 0
	regex, _ := regexp.Compile("mul\\(\\d+,\\d+\\)")

	instructions := regex.FindAllString(input, -1)

	for _, i := range instructions {
		temp := 0
		for index, v := range strings.Split(strings.ReplaceAll(string(i[4:]), ")", ""), ",") {
			switch index {
			case 0:
				temp += string_to_int(v)
			case 1:
				counter += temp * string_to_int(v)
			}
		}
	}

	fmt.Println("Part 1:", counter)
}

func part_2(input string) {
	counter := 0
	regex, _ := regexp.Compile("(do\\(\\))|(don't\\(\\))|mul\\(\\d+,\\d+\\)")
	enabled := true

	instructions := regex.FindAllString(input, -1)

	for _, i := range instructions {
		if i == "do()" {
			enabled = true
			continue
		} else if i == "don't()" {
			enabled = false
			continue
		} else {
			if !enabled { continue }

			temp := 0
			for index, v := range strings.Split(strings.ReplaceAll(string(i[4:]), ")", ""), ",") {

				switch index {
				case 0:
					temp += string_to_int(v)
				case 1:
					counter += temp * string_to_int(v)
				}
			}
		}
	}

	fmt.Println("Part 2:", counter)
}
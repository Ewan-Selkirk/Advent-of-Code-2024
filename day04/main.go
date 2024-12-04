package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"regexp"
	// "strings"
)

func main() {
	input := load_input()

// 	input := `MMMSXXMASM
// MSAMXMSMSA
// AMXSXMAAMM
// MSAMASMSMX
// XMASAMXAMM
// XXAMMXXAMA
// SMSMSASXSS
// SAXAMASAAA
// MAMMMXMMMM
// MXMXAXMASX`

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
	fields := strings.Fields(input)

	check_for_xmas := func(in_string string) bool {
		if in_string == "XMAS" || in_string == "SAMX" {
			counter++
			return true
		}

		return false
	}


	for line_index, line := range fields {
		for char_index, char := range line {
			actual_char := string(char)

			// God this is making me feel like YandereDev...
			if actual_char == "X" || actual_char == "S" {
				// Check Right
				if char_index <= len(line) - 4 { 
					char_range := line[char_index : char_index + 4]
					
					check_for_xmas(char_range)
				}

				// Check Left
				if char_index >= 3 {
					char_range := line[(3 - char_index) * -1 : char_index + 1]
					check_for_xmas(char_range)
				}

				// Check Up
				if line_index >= 3 {
					tmp_string := ""
					for i := 0; i < 4; i++ {
						tmp_string += string(fields[line_index - i][char_index])
					}

					check_for_xmas(tmp_string)
				}

				// Check Down
				if line_index <= len(fields) - 4 {
					tmp_string := ""
					for i := 0; i < 4; i++ {
						tmp_string += string(fields[line_index + i][char_index])
					}

					check_for_xmas(tmp_string)
				}

				// Check Diagonal Up Left
				if line_index >= 3 && char_index >= 3 {
					tmp_string := ""
					for i := 0; i < 4; i++ {
						tmp_string += string(fields[line_index - i][char_index - i])
					}

					check_for_xmas(tmp_string)
				}

				// Check Diagonal Up Right
				if line_index >= 3 && char_index <= len(line) - 4 {
					tmp_string := ""
					for i := 0; i < 4; i++ {
						tmp_string += string(fields[line_index - i][char_index + i])
					}

					check_for_xmas(tmp_string)
				}

				// Check Diagonal Down Left
				if line_index <= len(line) - 4 && char_index >= 3 {
					tmp_string := ""
					for i := 0; i < 4; i++ {
						tmp_string += string(fields[line_index + i][char_index - i])
					}

					check_for_xmas(tmp_string)
				}

				// Check Diagonal Down Right
				if line_index <= len(line) - 4 && char_index <= len(line) - 4 {
					tmp_string := ""
					for i := 0; i < 4; i++ {
						tmp_string += string(fields[line_index + i][char_index + i])
					}

					check_for_xmas(tmp_string)
				}
			}
		}
	}

	// I can't believe this worked... I suck at this shit lmao
	// Dividing by 2 feels like cheating but it worked so eh fight me
	fmt.Println("Part 1:", counter / 2)
}

func part_2(input string) {
	counter := 0
	fields := strings.Fields(input)

	// Regex my beloved
	regex, _ := regexp.Compile("(M\\wS\\wA\\wM\\wS)|(S\\wS\\wA\\wM\\wM)|(M\\wM\\wA\\wS\\wS)|(S\\wM\\wA\\wS\\wM)")

	for line_index := 0; line_index < len(fields) - 2; line_index++ {
		for char_index := 0; char_index < len(fields[line_index]) - 2; char_index++ {
			tmp_string := ""

			for i := 0; i < 3; i++ {
				tmp_string += fields[line_index + i][char_index : char_index + 3]
			}

			if regex.MatchString(tmp_string) { counter++ }
		}
	}

	fmt.Println("Part 2:", counter)
}
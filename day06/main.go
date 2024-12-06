package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type position struct {
	x int
	y int
}

var facing int = 0
var position_x int = 0
var start_x int = 0
var position_y int = 0
var start_y int = 0
var data []string
var visited []position
var exit bool = false

func main() {
	input := load_input()

	/*input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`*/

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

func handle_movement(counter *int) {
	check_pos := func(x int, y int) {
		if string(data[y][x]) == "#" {
			facing = (facing + 1) % 4
			handle_movement(counter)
		} else {
			if (x == len(data[0]) - 1 && facing == 1) || (x == 0 && facing == 3) || (y == len(data) - 1 && facing == 2) || (y == 0 && facing == 0) {
				*counter++
				
				fmt.Println("Part 1:", *counter)
				exit = true
				return
			}

			position_x = x
			position_y = y

			if slices.Contains(visited, position{x, y}) { return }

			*counter++
			visited = append(visited, position{x, y})
		}

		return
	}

	switch facing {
		// Up
	case 0:
		for y := position_y; y >= 0; y-- {
			check_pos(position_x, y)
			if exit { return }
		}
		// Right
	case 1:
		for x := position_x; x < len(data[0]); x++ {
			check_pos(x, position_y)
			if exit { return }
		}

		// Down
	case 2:
		for y := position_y; y < len(data); y++ {
			check_pos(position_x, y)
			if exit { return }
		}

		// Left
	case 3:
		for x := position_x; x >= 0; x-- {
			check_pos(x, position_y)
			if exit { return }
		}
	}
}

func check_for_loop(block_x int, block_y int, looped *int, counter *int) {
	check_pos := func(x int, y int, counter *int) bool {
		if string(data[y][x]) == "#" || (y == block_y && x == block_x) {
			facing = (facing + 1) % 4
			check_for_loop(block_x, block_y, looped, counter)
		} else {
			if *looped == 5 {
				*counter++
				exit = true
				return true
			} 

			if (position_x == len(data[0]) - 1 && facing == 1) || (position_x == 0 && facing == 3) || (position_y == len(data) - 1 && facing == 2) || (position_y == 0 && facing == 0) {
				exit = true
				return false
			} else if (position_x == start_x && position_y == start_y) {
				*looped++
			}

			position_x = x
			position_y = y
		}

		return false
	}

	switch facing {
		// Up
	case 0:
		for y := position_y; y >= 0; y-- {
			check_pos(position_x, y, counter)
			if exit { return }
		}
		// Right
	case 1:
		for x := position_x; x < len(data[0]); x++ {
			check_pos(x, position_y, counter)
			if exit { return }
		}

		// Down
	case 2:
		for y := position_y; y < len(data); y++ {
			check_pos(position_x, y, counter)
			if exit { return }
		}

		// Left
	case 3:
		for x := position_x; x >= 0; x-- {
			check_pos(x, position_y, counter)
			if exit { return }
		}
	}
}

func part_1(input string) {
	counter := 0
	data = strings.Fields(input)

	for line_index, line := range data {
		for char_index, char := range line {
			if string(char) == "^" {
				position_x = char_index
				position_y = line_index

				start_x = position_x
				start_y = position_y
				break
			}
		}
	}

	handle_movement(&counter)
}

// I give up with part 2 for now
// After staring at it for 4-ish hours I realised I was doing it wrong :(
// Maybe I'll come back to it so for now it's only part 1
func part_2(input string) {

}
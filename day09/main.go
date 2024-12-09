package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type block struct {
	id int
	file_block int
	free_block int
}

func main() {
	// input := load_input()

	input := "2333133121414131402"

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


func part_1(input string) {
	var blocks []block

	file := 0
	free := 0
	id := 0
	for index, char := range input {

		v, _ := strconv.Atoi(string(char))
		
		switch index % 2 {
		case 0:
			file = v

			if index == len(input) - 1 {
				blocks = append(blocks, block{id, file, free})
			}
		case 1:
			free = v

			blocks = append(blocks, block{id, file, free})
			file = 0
			free = 0
			id++
		}
	}

	var values []string
	for _, block := range blocks {
		for i := 0; i < block.file_block; i++ {
			values = append(values, strconv.Itoa(block.id))
		}

		for i := 0; i < block.free_block; i++ {
			values = append(values, ".")
		}
	}
	
	// Reorder
	for i, v := range values {
		if v != "." { continue }

		for rev := len(values) - 1; rev >= 0; rev-- {
			if values[rev] == "." { continue }
			if rev < i { break }

			values[i], values[rev] = values[rev], values[i]
			break
		}
	}

	checksum := 0
	for i, r := range values {
		if r == "." { break }

		value, _ := strconv.Atoi(r)
		checksum += i * value
	}

	fmt.Println("Part 1:", checksum)
}

func part_2(input string) {
	var blocks []block

	check_slice := func(slice []string, query string) bool {
		for _, v := range slice {
			if v != query { return false }
		}

		return true
	}

	file := 0
	free := 0
	id := 0
	for index, char := range input {

		v, _ := strconv.Atoi(string(char))
		
		switch index % 2 {
		case 0:
			file = v

			if index == len(input) - 1 {
				blocks = append(blocks, block{id, file, free})
			}
		case 1:
			free = v

			blocks = append(blocks, block{id, file, free})
			file = 0
			free = 0
			id++
		}
	}

	var values []string
	for _, block := range blocks {
		values = append(values, strings.Repeat(strconv.Itoa(block.id), block.file_block))

		for i := 0; i < block.free_block; i++ {
			values = append(values, ".")
		}
	}

	for i := len(values) - 1; i != 0; i-- {
		if values[i] == "." { continue }

		size := len(values[i])
		for idx, v := range values {
			if v != "." { continue }
			if idx > i { break }

			if !check_slice(values[idx : idx + size], ".") { continue }

			values[i], values[idx] = values[idx], values[i]
			if size - 1 == 0 { break }

			for j := 1; j < size; j++ {
				for k := idx + 1; k < i; k++ {
					values[k], values[k + 1] = values[k + 1], values[k]
				}
			}

			break
		}
	}

	// Create a string from array
	str := ""
	for _, v := range values {
		str += v
	}


	checksum := 0
	for i, r := range str {

		value, _ := strconv.Atoi(string(r))
		checksum += i * value
	}

	fmt.Println("Part 2:", checksum)
}

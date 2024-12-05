package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	// "regexp"
)

func main() {
	input := load_input()

/*	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
*/
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

func sort_by_rules(rules []string, input []int) int {
	valid := false

	for _, rule := range rules {
		value := string_to_int(strings.Split(rule, "|")[0])
		before := string_to_int(strings.Split(rule, "|")[1])

		if !(slices.Contains(input, value) && slices.Contains(input, before)) { continue }
		if !(slices.Index(input, value) < slices.Index(input, before)) { 
			sort.IntSlice.Swap(input, slices.Index(input, value), slices.Index(input, before))
			valid = true

			sort_by_rules(rules, input)
		}
	}

	if valid { return input[len(input) / 2] } else { return 0 }
}

func part_1(input string) {
	counter := 0
	rules := strings.Fields(strings.Split(input, "\n\n")[0])
	updates := strings.Fields(strings.Split(input, "\n\n")[1])

	check_rules := func(input []int) bool {
		for _, rule := range rules {
			value := string_to_int(strings.Split(rule, "|")[0])
			before := string_to_int(strings.Split(rule, "|")[1])

			if !(slices.Contains(input, value) && slices.Contains(input, before)) { continue }
			if slices.Index(input, value) < slices.Index(input, before) { 
				continue 
			} else {
				return false
			}
		}

		return true
	}

	for _, update_line := range updates {
		var tmp_list []int
		split_update := strings.Split(update_line, ",")

		for _, update := range split_update {
			tmp_list = append(tmp_list, string_to_int(update))
		}

		if check_rules(tmp_list) {
			counter += tmp_list[len(tmp_list) / 2]
		}
		 
	}

	fmt.Println("Part 1:", counter)
}

func part_2(input string) {
	counter := 0
	rules := strings.Fields(strings.Split(input, "\n\n")[0])
	updates := strings.Fields(strings.Split(input, "\n\n")[1])

	for _, update_line := range updates {
		var tmp_list []int
		split_update := strings.Split(update_line, ",")

		for _, update := range split_update {
			tmp_list = append(tmp_list, string_to_int(update))
		}

		counter += sort_by_rules(rules, tmp_list)
		 
	}

	fmt.Println("Part 2:", counter)
}
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Read from a text file
// Return a list of lines in string
func read_input(fp string) []string {
	dat, err := os.ReadFile(fp)
	check(err)
	s_li := strings.Split(strings.ReplaceAll(string(dat), "\r\n", "\n"), "\n")
	return s_li
}

// Convert a list of string to a list of integer
func convert_str_to_int_li(str_li []string) []int {
	int_li := make([]int, len(str_li))
	var err error
	for i, s := range str_li {
		int_li[i], err = strconv.Atoi(s)
		check(err)
	}
	return int_li
}

// (string, integer)
type action struct {
	act string
	num int
}

// Convert a list of string to a list of (string, integer)
func convert_str_to_str_int_li(str_li []string) []action {
	str_int_li := make([]action, len(str_li))
	var err error
	for i, s := range str_li {
		sep_item := strings.Split(s, " ")
		str_int_li[i].act = sep_item[0]
		str_int_li[i].num, err = strconv.Atoi(sep_item[1])
		check(err)
	}
	return str_int_li
}

func main() {
	s_li := read_input("../input.txt")
	s_i_li := convert_str_to_str_int_li(s_li)
	depth := 0
	h_pos := 0
	for _, s_i := range s_i_li {
		// fmt.Println(i, s_i)
		// break
		if s_i.act == "forward" {
			h_pos += s_i.num
		} else if s_i.act == "down" {
			depth += s_i.num
		} else if s_i.act == "up" {
			depth -= s_i.num
		}
	}
	// for i := 0; i < len(s_i_li); i++ {
	// 	fmt.Println(s_i_li[i])
	// 	break
	// }
	fmt.Println(depth * h_pos)
}

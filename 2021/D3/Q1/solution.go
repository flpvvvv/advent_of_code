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

// Convert a list of string to a list of lists of integer
func convert_str_to_li_int_li(str_li []string) [][]int {
	li_int_li := make([][]int, len(str_li))
	var err error
	for i, s := range str_li {
		li_int_li[i] = make([]int, len(s))
		for j, char := range s {
			// fmt.Println(i, j, string(char))
			// fmt.Println(i, j, char)
			li_int_li[i][j], err = strconv.Atoi(string(char))
			check(err)
		}
	}
	return li_int_li
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
	li_int_li := convert_str_to_li_int_li(s_li)
	// fmt.Println(li_int_li)
	var count_1 int
	total_digi := len(li_int_li)
	total_col := len(li_int_li[0])
	// gamma_li := make([]int, total_col)
	// epsilon_li := make([]int, total_col)
	gamma_str := make([]string, total_col)
	epsilon_str := make([]string, total_col)
	for j := 0; j < total_col; j++ {
		count_1 = 0
		for i := 0; i < total_digi; i++ {
			count_1 += li_int_li[i][j]
		}
		if count_1 > total_digi-count_1 {
			// gamma_li[j] = 1
			// epsilon_li[j] = 0
			gamma_str[j] = strconv.Itoa(1)
			epsilon_str[j] = strconv.Itoa(0)
		} else {
			// gamma_li[j] = 0
			// epsilon_li[j] = 1
			gamma_str[j] = strconv.Itoa(0)
			epsilon_str[j] = strconv.Itoa(1)
		}
	}
	// fmt.Println(gamma_li)
	// fmt.Println(epsilon_li)

	gamma_bi := strings.Join(gamma_str, "")
	epsilon_bi := strings.Join(epsilon_str, "")

	gamma, err := strconv.ParseInt(gamma_bi, 2, 64)
	check(err)
	epsilon, err := strconv.ParseInt(epsilon_bi, 2, 64)
	check(err)

	// fmt.Println(gamma)
	// fmt.Println(epsilon)

	fmt.Println(gamma * epsilon)
}

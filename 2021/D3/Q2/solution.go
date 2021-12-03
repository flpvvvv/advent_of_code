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

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func li_int_to_li_str(li_int []int) []string {
	li_str := make([]string, len(li_int))
	for i := 0; i < len(li_int); i++ {
		li_str[i] = strconv.Itoa(li_int[i])
	}
	return li_str
}

func main() {
	s_li := read_input("../input.txt")
	li_int_li := convert_str_to_li_int_li(s_li)
	// fmt.Println(li_int_li)
	var count_1 int
	var count_0 int
	total_digi := len(li_int_li)
	total_col := len(li_int_li[0])

	oxygen_li := make([]int, total_digi)
	CO2_li := make([]int, total_digi)

	for j := 0; j < total_digi; j++ {
		oxygen_li[j] = 1
		CO2_li[j] = 1
	}

	// gamma_str := make([]string, total_col)
	// epsilon_str := make([]string, total_col)
	for j := 0; j < total_col; j++ {
		count_1 = 0
		for i := 0; i < total_digi; i++ {
			count_1 += li_int_li[i][j] * oxygen_li[i]
		}
		count_0 = sum(oxygen_li) - count_1
		if count_1 >= count_0 {
			if count_0 == 0 {
				continue
			}
			for k := 0; k < total_digi; k++ {
				if li_int_li[k][j] == 0 {
					oxygen_li[k] = 0
				}
			}
		} else {
			if count_1 == 0 {
				continue
			}
			for k := 0; k < total_digi; k++ {
				if li_int_li[k][j] == 1 {
					oxygen_li[k] = 0
				}
			}
		}
		if sum(oxygen_li) == 1 {
			// fmt.Println(oxygen_li)
			break
		}
	}

	for j := 0; j < total_col; j++ {
		count_1 = 0
		for i := 0; i < total_digi; i++ {
			count_1 += li_int_li[i][j] * CO2_li[i]
		}
		count_0 = sum(CO2_li) - count_1
		if count_1 < count_0 {
			if count_1 == 0 {
				continue
			}
			for k := 0; k < total_digi; k++ {
				if li_int_li[k][j] == 0 {
					CO2_li[k] = 0
				}
			}
		} else {
			if count_0 == 0 {
				continue
			}
			for k := 0; k < total_digi; k++ {
				if li_int_li[k][j] == 1 {
					CO2_li[k] = 0
				}
			}
		}

		if sum(CO2_li) == 1 {
			// fmt.Println(CO2_li)
			break
		}
	}

	var oxygen, CO2 int64
	var err error
	for i := 0; i < total_digi; i++ {
		if oxygen_li[i] == 1 {
			oxygen, err = strconv.ParseInt(strings.Join(li_int_to_li_str(li_int_li[i]), ""), 2, 64)
			// fmt.Println(oxygen)
			check(err)
			break
		}
	}

	for i := 0; i < total_digi; i++ {
		if CO2_li[i] == 1 {
			// fmt.Println(li_int_li[i])
			CO2, err = strconv.ParseInt(strings.Join(li_int_to_li_str(li_int_li[i]), ""), 2, 64)
			// fmt.Println(CO2)
			check(err)
			break
		}
	}

	fmt.Println(oxygen * CO2)

}

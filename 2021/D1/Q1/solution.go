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

func read_input(fp string) []string {
	dat, err := os.ReadFile(fp)
	check(err)
	s_li := strings.Split(strings.ReplaceAll(string(dat), "\r\n", "\n"), "\n")
	return s_li
}

func convert_str_to_int_li(str_li []string) []int {
	int_li := make([]int, len(str_li))
	var err error
	for i, s := range str_li {
		int_li[i], err = strconv.Atoi(s)
		check(err)
	}
	return int_li
}

func main() {
	s_li := read_input("../input.txt")
	i_li := convert_str_to_int_li(s_li)
	count := 0
	for i := 1; i < len(i_li); i++ {
		if i_li[i] > i_li[i-1] {
			count++
		}
	}
	fmt.Println(count)
}

package main

import (
	"fmt"
)

func main() {

	s := ""
	fmt.Print(1<<32, "\n")
	fmt.Scan(&s)
	fmt.Print(myAtoi(s), "\n")

}

func myAtoi(s string) int {

	i, num, size := 0, 0, len(s)
	minFlag := false

	for ; i < size-1 && s[i] == ' '; i++ {
	}

	if i >= size {
		return num
	}

	switch s[i] {
	case '-':
		minFlag = true
		i++
	case '+':
		i++
	}

	for ; i < size; i++ {

		if s[i] < '0' || s[i] > '9' {
			break
		}

		num = num*10 + int(s[i]-'0')

		if minFlag && (num > Min32) {
			return -Min32
		}

		if !minFlag && num > Max32 {
			return Max32
		}

	}

	if minFlag {
		num = -num
	}

	return num

}

const (
	Min32 = (1 << 31)
	Max32 = (1 << 31) - 1
)

package main

import (
	"fmt"
)

func main() {

	s, rows := "", 0

	fmt.Scan(&s, &rows)
	fmt.Print(convert(s, rows), "\n")

}

func min(a int, b int) int {

	if a < b {
		return a
	}

	return b

}

func convert(s string, numRows int) string {

	ss := make([]string, numRows)
	sSize, tmpSize := len(s), 0
	ans := ""

	for i := 0; i < sSize; {

		tmpSize = min(numRows, sSize-i)

		for j := 0; j < tmpSize; j++ {

			ss[j] += string(s[i+j])

		}

		i += tmpSize

		if i >= sSize {
			break
		}

		tmpSize = min(numRows-2, sSize-i)

		for j := 0; j < tmpSize; j++ {

			ss[numRows-2-j] += string(s[i+j])

		}

		if tmpSize > 0 {
			i += tmpSize
		}
	}

	for i := 0; i < numRows; i++ {
		ans += ss[i]
	}

	return ans

}

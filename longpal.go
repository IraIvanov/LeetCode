package main

import (
	"fmt"
)

func main() {

	s := ""

	fmt.Scan(&s)
	//fmt.Print(s[0:1])
	fmt.Print(longestPalindrome(s), "\n")

}

func longestPalindrome(s string) string {

	size := len(s)

	for ; size > 1; size-- {

		fmt.Print(size, "\n")
		var ans string

		for i := 0; i <= (len(s) - size); i++ {

			ans = s[i : size+i]
			if isPal(ans) {
				return ans
			}

		}

	}

	return s[0:1]

}

func isPal(s string) bool {

	left, right := 0, len(s)-1

	for left < right {

		if s[left] != s[right] {
			return false
		}

		left++
		right--

	}

	return true

}

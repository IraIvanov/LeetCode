package main

import (
	"fmt"
)

func main() {

	var str string

	fmt.Scanf("%s", &str)
	ans := lengthOfLongestSubstring(str)

	//ans := test("aabaab!bb")
	fmt.Print(ans, "\n")

}

func lengthOfLongestSubstring(s string) int {

	charSet := make(map[byte]bool)
	maxSize, leftInd, i, size := 0, 0, 0, 0

	for i < len(s) {

		if !charSet[s[i]] {

			charSet[s[i]] = true
			i++
			size++

		} else {

			if size > maxSize {
				maxSize = size
			}

			repEl := s[i]

			for leftInd < i {

				charSet[s[leftInd]] = false
				size--
				leftInd++

				if !charSet[repEl] {
					break
				}

			}

		}

	}

	if size > maxSize {
		maxSize = size
	}

	return maxSize

}

package main

import (
	"fmt"
)

func main() {

	var num int
	var target int
	var nums []int
	fmt.Scanf("%d", &target)
	fmt.Scanf("%d", &num)

	for i := 0; i < num; i++ {

		var x int
		fmt.Scanf("%d", &x)
		nums = append(nums, x)

	}

	ans := twoSum(nums, target)
	fmt.Println(ans)

}

func twoSum(nums []int, target int) []int {

	var ans [2]int

	l := len(nums)

	for i := 0; i < l-1; i++ {

		tgt := target - nums[i]

		for j := i + 1; j < l; j++ {

			if nums[j] == tgt {

				ans[0], ans[1] = i, j
				return ans[:]

			}
		}

	}

	return ans[:]

}

package main

import (
	"fmt"
)

func main() {

	var num int
	var target int
	fmt.Scanf("%d", &target)
	fmt.Scanf("%d", &num)
	nums := make([]int, num)

	for i := 0; i < num; i++ {

		fmt.Scanf("%d", &(nums[i]))

	}

	ans := twoSum(nums, target)
	fmt.Println(ans)

}

func twoSum(nums []int, target int) []int {

	var ans [2]int
	seen := make(map[int]int)

	l := len(nums)

	for i := 0; i < l; i++ {

		tgt := target - nums[i]

		if j, ok := seen[tgt]; ok && j != i {

			ans[0], ans[1] = i, j
			return ans[:]

		}

		seen[nums[i]] = i

	}

	return ans[:]

}

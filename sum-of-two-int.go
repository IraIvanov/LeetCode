package main

import (
	"fmt"
)

func main() {

	var a, b int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println(getSum(a, b))

}

func getSum(a, b int) int {

	and := a & b
	xor := a ^ b

	if and == 0 {
		return xor
	}

	and = and << 1

	return getSum(and, xor)

}

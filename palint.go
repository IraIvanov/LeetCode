package main

import "fmt"

func main() {

	num := -121
	fmt.Print(isPal(num))
}

func isPal(num int) bool {

	reversed, x := 0, num

	for ; x > 0; x /= 10 {

		reversed = reversed*10 + x%10

	}

	return reversed == num

}

/*

Problem Statement

Convert the number to string

*/

// Solution

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int = 49
	fmt.Println("Value of number variable: ", number)
	fmt.Printf("Type of number variable: %T\n", number)
	str := strconv.Itoa(number)
	fmt.Println("Value of str variable: ", str)
	fmt.Printf("Type of str variable: %T\n", str)
}

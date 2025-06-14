/*

Problem Statement

Swap two numbers without using a third variable

*/

// Solution

package main

import "fmt"

func main() {
	var num1 int = 49
	var num2 int = 36
	fmt.Printf("value of num1: %d and num2: %d\n", num1, num2)
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Printf("Value of num1: %d and num2: %d\n", num1, num2)
}

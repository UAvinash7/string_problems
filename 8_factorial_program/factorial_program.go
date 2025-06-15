/*

Problem Statement

Factorial Program

*/

// Solution

package main

import "fmt"

func main() {
	var num int = 5	// calculating the value for factorial of 5
	var result int = 1
	for i := 1; i <= num; i++ {
		result = result  * i
	}
	fmt.Println("result:", result)
}

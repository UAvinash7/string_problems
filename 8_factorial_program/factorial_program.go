/*

Problem Statement

Factorial Program

*/

package main

import "fmt"

/*
func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}
*/

func main() {
	var num int = 5
	var res int = 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	fmt.Println("result:", res)

}

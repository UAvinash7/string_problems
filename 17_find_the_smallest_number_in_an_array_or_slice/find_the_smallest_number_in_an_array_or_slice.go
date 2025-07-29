/*

Problem Statement

Find the smallest number in an array or slice.

*/

// Solution

package main

import "fmt"

func main() {
	slice := []int{2, 4, 6, 1, 7}
	min_value := slice[0]
	for i := 0; i < len(slice); i++ {
		if min_value > slice[i] {
			min_value = slice[i]
		}
	}
	fmt.Println("Value of smallest Element:", min_value)
}

/*

Problem Statement

Search an element in the slice

*/

// Solution

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	var found int = 0
	var num int = 5
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			found = 1
		}
	}
	if found == 1 {
		fmt.Printf("Element %d is present in the array arr\n", num)
	} else {
		fmt.Printf("Elemetn %d is not present in the array arr\n", num)
	}
}

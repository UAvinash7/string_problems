/*

Problem Statement

Find the highest number in an array or slice

*/

// Solution

package main

import "fmt"

func main() {
	slice := []int{1, 4, 2, 6, 5, 3}
	max_value := slice[0]
	for i := 0; i < len(slice); i++ {
		if slice[i] > max_value {
			max_value = slice[i]
		}
	}
	fmt.Println("maximum value of element in the slice:", max_value)
}

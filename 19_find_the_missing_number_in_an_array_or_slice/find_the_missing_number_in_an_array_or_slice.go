/*

Problem Statement

Find the missing number in an array or slice.

*/

// Solution

package main

import "fmt"

func main() {
	slice := []int{2, 4, 1, 5, 7, 8}
	num := len(slice)
	total := (num + 1) * (num + 2) / 2
	for i := 0; i < num; i++ {
		total = total - slice[i]
	}
	fmt.Println("missing number:", total)
}

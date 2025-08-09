/*

Problem Statement

Lower to higher numbers and higher to lower numbers using a slice.

*/

// Solution

package main

import "fmt"

func main() {
	slice := []int{2, 1, 3, 5, 4, 7, 6}
	for i := 0; i < len(slice); i++ {
		for k := 0; k < len(slice)-1; k++ {
			if slice[k] > slice[k+1] {
				temp := slice[k+1]
				slice[k+1] = slice[k]
				slice[k] = temp
			}
		}
	}
	fmt.Println("slice:", slice)
}

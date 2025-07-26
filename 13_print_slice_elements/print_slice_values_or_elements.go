/*

Problem Statement

Print a slice (array) values or elements

*/

// Solution

package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 4, 2}
	fmt.Println("arr:", arr)
	for _, ele := range arr {
		if ele == 2 {
			fmt.Println("My ID is:", ele)
		}
	}
}

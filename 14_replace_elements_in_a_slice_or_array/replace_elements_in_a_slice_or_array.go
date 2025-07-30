/*

Problem Statement

Replace elements in a slice or array

*/

// Solution

package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{6, 7, 8}
	fmt.Println("arr1:", arr1)
	copy(arr1[3:], arr2)
	fmt.Println("new array:", arr1)
	fmt.Println("second array:", arr2)
}

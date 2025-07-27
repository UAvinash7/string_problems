/*

Problem Statement

Convert an array into a slice

*/

// Solution

package main

import "fmt"

func main() {
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := []int{}
	fmt.Println("Value of slice variable arrr2 is:", arr2)
	arr2 = arr1[:]
	fmt.Println("After assigning, the value of arr2 is:", arr2)
	arr2 = append(arr2, 5)
	fmt.Println("After appending, the value of arr2 is:", arr2)
}

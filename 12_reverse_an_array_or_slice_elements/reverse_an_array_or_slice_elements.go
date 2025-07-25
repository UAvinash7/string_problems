/*

Problem Statement

Reverse an array or slice elements

*/

// Solution

package main

import "fmt"

func main() {
	arr := []int{3, 4, 2, 7, 1, 5}
	fmt.Println("Value of arr:", arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println("Reverse of arr:", arr)
}

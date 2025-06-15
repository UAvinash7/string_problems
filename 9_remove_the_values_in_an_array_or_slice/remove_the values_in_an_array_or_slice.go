/*

Problem Statement

Remove the values in an array or slice

*/

// Solution

package main

import "fmt"

func main() {
	arr := []string{"one", "two", "three", "four"}
	fmt.Println("arr:", arr)		// arr: [one two three four]
	arr1 := []string{}
	fmt.Println("arr1:", arr1)		// arr1: []
	for i := 0; i < len(arr); i++ {
		arr1 = append(arr1, arr[i])
	}
	fmt.Println("arr1:", arr1)		// arr1: [one two three four]
	var arr2 []string
	fmt.Println("arr2:", arr2)		// arr2: []
	for i := 0; i < len(arr); i++ {
		if i == 2 {
			continue
		}
		arr2 = append(arr2, arr[i])		// arr2: [one two four]
	}
	fmt.Println("arr2:", arr2)
	var arr3 []string
	arr3 = arr[1:3]
	fmt.Println("arr3:", arr3)			// arr3: [two three]
	arr4 := []string{}
	arr4 = arr[:2]
	fmt.Println("arr4", arr4)			// arr4: [one two]
	arr5 := []string{}
	arr5 = arr[:]
	fmt.Println("arr5:", arr5)			// arr5: [one two three four]
}

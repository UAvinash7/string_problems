/*

Problem Statement

Count duplicate values in an array or slice

*/

// Solution

package main

import "fmt"

func main() {
	arr := []int{2, 3, 1, 4, 3, 2, 5, 2}
	count := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if count[arr[i]] < 1 {
			count[arr[i]] = 1
		} else {
			count[arr[i]] += 1
		}
	}
	for index, element := range count {
		fmt.Println("my number:", index, "and count:", element)
	}
}

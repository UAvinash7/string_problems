/*

Problem Statement

Uncommon characters in two strings

*/

// Solution

package main

import "fmt"

func main() {
	str1 := "afbde"
	str2 := "wabq"
	mpp := make(map[string]int)
	for i := 0; i < len(str1); i++ {
		mpp[string(str1[i])]++
	}
	for i := 0; i < len(str2); i++ {
		mpp[string(str2[i])]++
	}
	var result string
	for index, element := range mpp {
		if element <= 1 {
			result = result + index
		}
	}
	fmt.Println("result:", result)
}

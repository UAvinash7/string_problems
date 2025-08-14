/*

Problem Statement

Check if strings are rotation of each other or not

*/

// Solution

package main

import "fmt"

func main() {
	str1 := "abcde"
	str2 := "deabc"
	mpp := make(map[string]int)
	if len(str1) != len(str2) {
		fmt.Println("Second string is not rotation of first string")
	}
	for i := 0; i < len(str1); i++ {
		mpp[string(str1[i])]++
	}
	for i := 0; i < len(str2); i++ {
		mpp[string(str2[i])]--
	}
	result := 1
	for ind, _ := range mpp {
		if mpp[ind] != 0 {
			result = 0
		}
	}
	if result != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

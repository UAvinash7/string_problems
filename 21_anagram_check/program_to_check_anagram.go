/*

Problem Statement

Program to check anagram

*/

// Solution

package main

import "fmt"

func main() {
	str1 := "anagram"
	str2 := "nagaram"
	mpp := make(map[string]int)
	for i := 0; i < len(str1); i++ {
		mpp[string(str1[i])]++
	}
	for i := 0; i < len(str2); i++ {
		mpp[string(str2[i])]--
	}
	result := 1
	for i := 0; i < len(str1); i++ {
		if mpp[string(str1[i])] != 0 {
			result = 0
		}
	}
	if result == 1 {
		fmt.Println("Yes, this is an anagram")
	} else {
		fmt.Println("No, this is not an anagram")
	}
}

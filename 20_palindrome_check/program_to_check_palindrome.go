/*

Problem Statement

Program to check palindrome

*/

// Solution

package main

import "fmt"

func main() {
	str := "ABBA"
	result := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	if str == string(result) {
		fmt.Println("This is palindrome")
	} else {
		fmt.Println("This is not palindrome")
	}
}

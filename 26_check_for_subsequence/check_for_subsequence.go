/*

Problem Statement

Check for subsequence

*/

// Solution

package main

import "fmt"

func main() {
	str1 := "gksrek"
	str2 := "geeksforgeeks"
	n1 := len(str1)
	n2 := len(str2)
	t := check(str1, str2, n1, n2)
	if t == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func check(str1 string, str2 string, n1 int, n2 int) bool {
	if n1 == 0 {
		return true
	}
	if n2 == 0 {
		return false
	}
	if str1[n1-1] == str2[n2-1] {
		return check(str1, str2, n1-1, n2-1)
	}
	return check(str1, str2, n1, n2-1)
}

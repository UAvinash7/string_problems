/*

Problem Statememt

Alternatively merge two strings.

*/

// Solution

package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "Bye"
	// Output =  HBeylelo
	var result string
	for i := 0; i < len(str1) || i < len(str2); i++ {
		if i < len(str1) {
			result += string(str1[i])
		}
		if i < len(str2) {
			result += string(str2[i])
		}
	}
	fmt.Println("Result:", result)
}

/*

Problem Statement

Replace the word in a string

*/

// Solution

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello World!, hello, Universe!, hello Galaxy!"
	replaceStr := strings.Replace(str, "hello", "Hello", 3)
	// replaceStr := strings.Replace(str, "h", "H", 3)	// This also works
	fmt.Println("New string value: ", replaceStr)
}

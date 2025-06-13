/*

Problem Statement

Convert lower case string to upper case string and vice versa

*/

// Solution

package main

import (
	"fmt"
	"strings"
)

func main() {
	strLower := "hello world!"
	fmt.Printf("converting lower case string %s into upper case string: %s\n", strLower, strings.ToUpper(strLower))
	strUpper := "HELLO WORLD!"
	fmt.Printf("converting upper case string %s into lower case string: %s\n", strUpper, strings.ToLower(strUpper))
}

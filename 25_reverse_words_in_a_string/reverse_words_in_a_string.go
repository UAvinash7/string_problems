/*

Problem Statement

Reverse words in a string.

*/

// Solution

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "mod.pkg"
	fmt.Println("str:", str)
	arr := strings.Split(str, ".")
	var result string
	for _, ele := range arr {
		result = "." + ele + result
	}
	result = strings.Replace(result, ".", "", 1)
	fmt.Println("result:", result)
}

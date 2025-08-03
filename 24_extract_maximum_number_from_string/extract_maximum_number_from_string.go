/*

Problem Statement

Extract maximum number from string

*/

// Solution

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "100ab546sc200" // output = 546
	var result string
	for _, ele := range str {
		if string(ele) < "a" || string(ele) > "z" {
			result += string(ele)
		} else {
			result += " "
		}
	}
	arr := strings.Split(result, " ")
	num := "0"
	for i := 0; i < len(arr); i++ {
		if num < arr[i] {
			num = arr[i]
		}
	}
	fmt.Println("num:", num)
}

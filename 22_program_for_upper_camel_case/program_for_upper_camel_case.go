/*

Problem Statement

Program for upper camel case

*/

// Solution

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello ram" // Hello Ram
	fmt.Println("str:", str)
	slice := strings.Split(str, " ")
	fmt.Println("slice:", slice)
	var str1 string
	for _, ele := range slice {
		bs := []byte(ele)
		bs[0] = byte(bs[0] - 32)
		str1 = str1 + string(bs) + " "
	}
	fmt.Println("str1:", str1)

}

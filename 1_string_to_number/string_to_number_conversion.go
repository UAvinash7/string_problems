/*

Problem Statement

Convert the string to number

*/

// Solution

package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100"
	fmt.Println("value of str variable:", str)
	fmt.Printf("str variable is of type: %T\n", str)
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("value of num variable:", num)
	fmt.Printf("num variable is of type: %T\n", num)
}
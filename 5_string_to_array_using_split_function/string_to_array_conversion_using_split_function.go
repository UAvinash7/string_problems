/*

Problem Statement

Convert the string into an array using split function

*/

// Solution

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "I love my country"
	fmt.Println("str:", str)
	arr := strings.Split(str, " ")
	fmt.Println("arr:", arr)
	for index, element := range arr {
		fmt.Printf("index: %d, element: %s\n", index, element)
	}
}

/*

func main() {
	var str string = "accomodation"
	arr := strings.Split(str, "")
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
}

*/
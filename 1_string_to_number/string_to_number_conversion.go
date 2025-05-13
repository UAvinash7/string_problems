// Array and Slice
// a := [2]int32{} and b := [3]int32{} are of different types

// arrays are not growable, you cannot append more elements to an existing array.

// Slice is a data  structure describing a continguous section of an array.
// Slice is not an array.

// Slice is represented in memory with three elements :
// length, capacity and pointer to the first element in the backing array.

// the pointer part will be having the address of the 5th element of the numbers array since we
// have specified index 4 at the start of the slice

// The second represents the length of the array with is 2 in this case.

// The third is the capacity of the backing array.

// The original numbered array has the capacity of 6 elements.
// However our slice starts from index 4 only and thus it can use only the storage of the numbers array
// starting from that location. Thus, the capacity will be 2

package main

import "fmt"

/*
func main() {
	avi := []int32{0, 1}
	helper(avi)
}

func helper(a []int32) {
	fmt.Println(a)
}

*/

func main() {
	var numbers = [6]int{81, 64, 49, 36, 25, 16}
	var slice []int = numbers[2:4] //	[49, 36]
	fmt.Println("slice:", slice)
	fmt.Println("length of slice:", len(slice))
	fmt.Println("capacity of slice:", cap(slice))
	fmt.Println("slice[2]:", slice[1])
	slice = slice[0:4]
	fmt.Println("slice:", slice)
	slice = slice[:]
	
}

// Array and Slice
// a := [2]int32{} and b := [3]int32{} are of different types

// arrays are not growable, you cannot append more elements to an existing array.

// Slice is a data structure describing a continguous section of an array.
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

func substractOneFromLength(slice *[]int) *[]int{
	newSlice := *slice
	newSlice = newSlice[0 : len(newSlice)-1]
	*slice = newSlice
	return slice
}

func main() {
	avi := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("avi:", avi)		// [0 1 2 3 4 5]
	fmt.Println("avi's length:", len(avi))	// 6
	fmt.Println("avi's capacity:", cap(avi))	// 6
	res := substractOneFromLength(&avi)
	fmt.Println("after function call")
	fmt.Println("avi:", avi)	// [0 1 2 3 4 5]
	fmt.Println("avi's length:", len(avi))	// 6
	fmt.Println("avi's capacity:", cap(avi))	// 6
	fmt.Println("res:", *res)	// [0 1 2 3 4]
	fmt.Println("length of res:", len(*res))	// 5
	fmt.Println("capacity of res:", cap(*res))	// 6

}

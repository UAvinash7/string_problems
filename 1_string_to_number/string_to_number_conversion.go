// Array and Slice
// a := [2]int32{} and b := [3]int32{} are of different types

// arrays are not growa  ble, you cannot append more elements to an existing array.

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

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	message := Message{
		Title: "Hello World",
		Body:  "This is a wonderful message, I am sending to you",
	}
	data, err := json.Marshal(message)
	checkError(err)
	fmt.Println(string(data))
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.Encode(message)
	//file, err := os.Create("message.json")
	//checkError(err)
	//defer file.Close()
	//io.Copy(file, buf)
	fmt.Println(buf.String())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

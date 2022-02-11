package main

import (
	"fmt"
	"math"
	"strings"
)

func no_longer_main() {

	// POINTER
	// pointers hold the memory address of a value
	var pointer *int
	pointer_value := 42
	pointer = &pointer_value
	fmt.Printf("Value: %v\tStored at: %v\n", *pointer, pointer)


	// STRUCT
	// A struct is a collection of fields.
	// (see Vertex)
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)

	struct_pointer := &v
	struct_pointer.X = 1e9 // use this notation instead of (*struct_pointer).X
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)

	// ARRAYS
	// cannot be resized
	var my_array [10]string
	my_array[0] = "Hello"
	my_array[1] = "World"
	fmt.Println(my_array[0], my_array[1])
	fmt.Println(my_array)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// SLICES 
	/*
	An array has a fixed size. 
	A slice, on the other hand, 
	is a dynamically-sized, flexible 
	view into the elements of an array. 
	In practice, slices are much more 
	common than arrays.
	*/

	var primes_slice [] int = primes[1:4] // first/last indices can be omitted (e.g. [:10] or [4:])
	fmt.Println(primes_slice)

	// Changing the elements of a slice 
	// modifies the corresponding elements 
	// of its underlying array
	primes_slice[0] = 123
	fmt.Println(primes_slice)
	fmt.Println(primes)

	// slice literals
	slice_literal := []bool{true, true, false}
	fmt.Println(slice_literal)

	// array/slice info
	fmt.Printf("primes info:\n\tlength=%d\n\tcapacity=%d\n\tcontents=%v\n", len(primes), cap(primes), primes)
	fmt.Printf("primes_slice info:\n\tlength=%d\n\tcapacity=%d\n\tcontents=%v\n", len(primes_slice), cap(primes_slice), primes_slice)
	primes_slice = primes[2:]
	fmt.Printf("primes_slice info:\n\tlength=%d\n\tcapacity=%d\n\tcontents=%v\n", len(primes_slice), cap(primes_slice), primes_slice)


	var empty_slice []int
	fmt.Println(empty_slice, len(empty_slice), cap(empty_slice))
	if empty_slice == nil {
		fmt.Println("nil!")
	}

	make_slice_1 := make([]int, 0, 5)
	printIntSlice("make_slice_1", make_slice_1)

	make_slice_2 := make([]int, 5) // can exlcude capacity arg
	printIntSlice("make_slice_2", make_slice_2)

	// 2-d slices
	// create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// appending to a slice
	var dyanamic_slice []int
	printIntSlice("Dynamic Slice", dyanamic_slice)
	dyanamic_slice = append(dyanamic_slice, 1, 2, 3)
	printIntSlice("Updated Dynamic Slice", dyanamic_slice)
	// note the automatic capacity allocation


	// range form of the for loop
	var pow = []int{1, 2, 3, 4, 5}
	get_squares(pow)

	nums_slice := []int{0, 1, 2, 3, 4}
	for i := range nums_slice {
		nums_slice[i] += 1
	}
	printIntSlice("nums_slice", nums_slice)

	// MAPS (think Python dictionary)
	coordinates_map := make(map[string]Coordinates)
	coordinates_map["Bell Labs"] = Coordinates{40.68433, -74.39967,}
	fmt.Println(coordinates_map["Bell Labs"])

	// map literals
	google_map := map[string]Coordinates{
		"Google": {37.42202, -122.08408,
		},
	}
	fmt.Println(google_map)

	// mutating maps
	google_map["New HQ"] = Coordinates{0, 0}
	delete(google_map, "New HQ")

	value, present := google_map["New HQ"]
	fmt.Println("Value: ", value, "\tPresent? ", present)

	value, present = google_map["Google"]
	fmt.Println("Value: ", value, "\tPresent? ", present)


	/* 
	Functions are values too. They can be passed around just like other values.
	Function values may be used as function arguments and return values.
	*/

}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}	// has type Vertex
	v2 = Vertex{X: 1}	// Y: 0  is implicit
	v3 = Vertex{} 		// X: 0 and Y: 0
	p = &Vertex{1, 2}	// has type *Vertex
)

func printIntSlice (s string, a []int) {
	fmt.Printf("%s Info:\n\tlength=%d\n\tcapacity=%d\n\tcontents=%v\n", s, len(a), cap(a), a)
} 

func get_squares(a []int) {
	// two values are returned here: index and element
	// you can skip i or v by typing _ instead (e.g. for _, v := range a{})
	// if you want only the index, you can omit the second var
	for i, v := range a {
		fmt.Printf("Index: %v\t%d**2 = %d\n", i, v, int(math.Pow(float64(v), 2)))
	}
}

type Coordinates struct {
	Lat, Long float64
}
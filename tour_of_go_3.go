package main

import (
	"fmt"
	"time"
)

func hitherto_not_main() {

	// SLICES AND MAPS PRACTICE
	var numbers = []int{0, 1, 2, 3, 4, 5}	// capacity set to 6
	printIntSlice("numbers", numbers)
	numbers = append(numbers, 6)			// capacity now set to 12 
	printIntSlice("numbers", numbers)
	numbers = numbers[1:]
	printIntSlice("numbers", numbers)		// capacity now set to 11

	grades := make(map[string]int)
	grades["student 1"] = 95
	grades["student 2"] = 90
	grades["student 3"] = 85
	grades["student 4"] = 80
	fmt.Println(grades)
	fmt.Println("Map Length:", len(grades))
	delete(grades, "student 4")
	fmt.Println(grades)
	fmt.Println("Map Length:", len(grades))

	// USER INPUT
	var first string
	var last string

	fmt.Print("First name: ")
	fmt.Scanln(&first)

	fmt.Print("Last name: ")
	fmt.Scanln(&last)

	fmt.Println("Hello,", first, last + "!")

	// SLEEP function
	fmt.Print("Loading")
	wait_a_sec()
	fmt.Print(".")
	wait_a_sec()
	fmt.Print(".")
	wait_a_sec()
	fmt.Print(".")
	wait_a_sec()
	fmt.Println("\nLoading Complete!")
	wait_a_sec()
}

func wait_a_sec() {
	time.Sleep(1 * time.Second)
}

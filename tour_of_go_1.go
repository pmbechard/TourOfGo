package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func not_main_anymore() {

	// Variable assignments
	var yes, no, maybe bool
	var i int

	yes, no, maybe = true, false, true
	i = 5
	var j, k int = 6, 7
	l := 8 

	/*
	Variables declared without an explicit initial value are given their zero value.
	i.e. 0, "", false
	*/

	fmt.Println(yes, no, maybe, i, j, k, l)

	var message = "Hello, world!"
	fmt.Println(message)
	fmt.Println("Hello, string literal!")

	message = "Hey, there!"
	fmt.Println(message)

	/*
	The expression T(v) converts the value v to the type T.

		Some numeric conversions:

		var i int = 42
		var f float64 = float64(i)
		var u uint = uint(f)
		
		Or, put more simply:

		i := 42
		f := float64(i)
		u := uint(f)
	*/

	int_number := int(25)
	fmt.Printf("Type: %T, Value: %v\n", int_number, int_number)
	float_number := float64(58.4)
	fmt.Printf("Type: %T, Value: %v\n", float_number, float_number)
	complex_number := complex128(4.006893)
	fmt.Printf("Type: %T, Value: %v\n", complex_number, complex_number)

	var number2 = 25
	var sum = int_number + number2
	fmt.Println(sum)

	// Constants
	const Pi = 3.14
	// Pi += 1 causes an error


	// shift decimal places 10 right??
	const shifty_number = 1 << 10
	fmt.Println(shifty_number)




	var dictionary = make(map[string]int)
	dictionary["one"] = 1
	dictionary["two"] = 2
	dictionary["three"] = 3
	fmt.Println(dictionary)
	fmt.Println(dictionary["two"])

	var name = "Peyton"
	var age = 29
	fmt.Printf("Hello, my name is %v. I am %v years old.\n", name, age)
	age_string := strconv.Itoa(age)
	
	fmt.Println("Hello, my name is " + name + ". I am " + age_string + " years old.")

	age_reint, err := strconv.Atoi(age_string)

	if err == nil {
		fmt.Println(age_reint + 1)
	}


	var username string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Printf("Welcome back, %v!\n", username)

	var user_age int
	fmt.Print("Age: ")
	fmt.Scanln(&user_age)
	fmt.Printf("You are %v years old.\n", age)

	fmt.Println("My favourite number is", rand.Intn(10))
	var two_up_three = math.Pow(2, 3)
	fmt.Println(two_up_three)

	var array [10]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	var add_method_sum = add(3, 4)
	fmt.Println("Sum: ", add_method_sum)

	fmt.Println(swap("hello", "world"))

	fmt.Println(split(100))
	fmt.Println(split(17))

	describe_int_var(i)
	describe_int_var(j)
	describe_int_var(k)
	describe_int_var(l)

	var loop_sum int = 0

	for i := 0; i < 10; i++ {
		loop_sum += i
	}
	fmt.Println(loop_sum)

	// shorthand for loop (init and post statements are optional)
	loop_sum = 1
	for ; loop_sum < 1000; {
		loop_sum += loop_sum
	}
	fmt.Println(loop_sum)

	// while equivalent
	loop_sum = 1
	for loop_sum < 1000 {
		loop_sum += loop_sum
	}
	fmt.Println(loop_sum)

	// infinite loop
	// for {
	// }

	if loop_sum >= 1000 {
		fmt.Println("The var is over 1000.")
	} else if loop_sum >= 500 {
		fmt.Println("The var is over 500.")
	} else {
		fmt.Println("The var is under 500.")
	}

	// if statements can start with a command to execute before the condition
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 20),
	)

	// switch statements
	check_system()
	// unlike Java, stops when a case succeeds

	// runs as switch true {} for long if-then-else chains
	// switch {
	// }

	// defer statement
	get_time()

	// if there are multiple defers, they are pushed onto a LIFO stack
	fmt.Println("\n\nCounting...")
	for i:= 0; i <= 10; i++ {
		defer fmt.Println(i)
	}  
	fmt.Println("Done.") // printed before numbers.

}

// if two or more consecutive params are the same type, omit all but the last to shorten the code 
func add(x, y int) int {
	return x + y
}

// functions can return any number of values
func swap(a, b string) (string, string) {
	return b, a
}

// return named variables
func split(sum int) (x, y int) {
	if sum % 2 == 0 {
		x = sum / 2
		y = sum / 2
	} else {
		x = sum / 2
		y = sum - x
	}
	return
}

func describe_int_var(x int) {
	fmt.Printf("Type: %T     Value: %v\n", x, x)
}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func check_system() {
	fmt.Println("Running on...")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("\t-OS X")
	case "linux":
		fmt.Println("\t-Linux")
	default:
		fmt.Printf("\t-%s \n", os)
	}
}

func get_time() {
	defer fmt.Println("Function complete.")
	fmt.Println(time.Now().Clock())
}

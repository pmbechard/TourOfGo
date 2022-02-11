package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	// CONCURRENCY

	// GOROUTINES

	// Example
	go count("go keyword")
	time.Sleep(time.Millisecond * 200)
	count("regular method call")
	/* 
	goroutine will exit with end of home function
	(main in this case), so if "regular mathod call"
	finishes first, goroutine will end at that point
	whether or not it has finished

	See the solution to this issue below:
	*/ 


	fmt.Println("\nPress enter to continue...")
	fmt.Scanln()


	// sync.WaitGroup
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	go func() {
		count("Synced")
		waitgroup.Done()
	}()

	waitgroup.Wait()
	// The above block will ensure that the   
	// goroutine has finished before proceeding


	fmt.Println("\nPress enter to continue...")
	fmt.Scanln()


	// CHANNELS
	c := make(chan string)
	go count_with_channels("Using Channels", c)

	for message := range c{
		fmt.Println(message)
	}

	
	fmt.Println("\nPress enter to continue...")
	fmt.Scanln()


	// Channel blockages
	// when a channel is waiting for a messages the current
	// routine stops

	/*
	e.g.

	c: = make(chan string)
	c <- "hello"				<-- code stops here to wait
	
	msg = <- c
	fmt.Println(msg)
	*/

	// Buffered Channel (add buffer cap as an arg)
	buffered_channel := make(chan string, 2)
	buffered_channel <- "hello"
	buffered_channel <- "world"

	msg := <- buffered_channel
	fmt.Println(msg)

	msg = <- buffered_channel
	fmt.Println(msg)
	
	
	fmt.Println("\nPress enter to continue...")
	fmt.Scanln()


	selected_channel_1 := make(chan string)
	selected_channel_2 := make(chan string)

	go func() {
		for i := 0; i < 20; i++ {
			selected_channel_1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		} 
	}()

	go func() {
		for i := 0; i < 5; i++ {
			selected_channel_2 <- "Every 2s"
			time.Sleep(time.Second * 2)
		}
	}()

	// Displays syncronous output - ends in deadlock error
	// for i := 0; i < 40; i++ {
	// 	select {
	// 	case msg1 := <- selected_channel_1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <- selected_channel_2:
	// 		fmt.Println(msg2)
	// 	}
	// }

	// Worker Pools
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	// single worker is very slow
	// 4 workers working concurrently is faster
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	// However, there is no guarantee the results
	// will appear in order...

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(j, <-results)
	}

}

func count(s string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}

func count_with_channels(s string, c chan string) {
	for i := 1; i <= 5; i++ {
		result := []string{s, strconv.Itoa(i)}
		c <- strings.Join(result, " - ")
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

// <-s represent that we will only ever send on the jobs channel
// and only ever receive on the results channel
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}
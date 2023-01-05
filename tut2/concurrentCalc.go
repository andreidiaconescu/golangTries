package main

import "fmt"

var count = 45

func main() {
	chJobs := make(chan int)
	chResults := make(chan int)

	go workerFib(chJobs, chResults)
	go workerFib(chJobs, chResults)
	go workerFib(chJobs, chResults)
	go workerFib(chJobs, chResults)

	go func() {
		for i := 0; i < count; i++ {
			chJobs <- i
		}
		close(chJobs)
	}()

	for j := 0; j < count; j++ {
		fmt.Println(<-chResults)
	}
}

func workerFib(chFibN <-chan int, chResults chan<- int) {
	for n := range chFibN {
		chResults <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

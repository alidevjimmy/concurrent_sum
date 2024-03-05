package v2

import (
	"log"
	"runtime"
	"sync"
)

func concurrentSum(arr []int) int {
	// get number of logical CPUs usable by the process
	coreCount := runtime.NumCPU()

	// calculate array chunk size
	chunkSize := len(arr) / coreCount
	if coreCount > chunkSize {
		chunkSize = 1
	}
	// create chunks from array and calculate
	// sum of each array in corresponding goroutine
	var wg sync.WaitGroup

	res := make(chan int, coreCount+1)

	for i := 0; i < len(arr); i += chunkSize {
		wg.Add(1)

		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		go calculate(arr[i:end], &wg, res)
	}

	wg.Wait()
	close(res)

	sum := 0
	for r := range res {
		sum += r
	}

	return sum
}

func calculate(arr []int, wg *sync.WaitGroup, res chan<- int) {
	sum := sumOfArray(arr)

	defer wg.Done()

	select {
	case res <- sum:
	default:
		log.Printf("error: cannot send sum to res channel")
	}

}

func sumOfArray(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func naiveConcurrentSum(arr []int) int {
	return sumOfArray(arr)
}

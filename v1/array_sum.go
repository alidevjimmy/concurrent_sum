package v1

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

	var sums []<-chan int

	for i := 0; i < len(arr); i += chunkSize {
		wg.Add(1)

		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		sums = append(sums, calculate(arr[i:end], &wg))
	}

	wg.Wait()

	res := 0
	for sum := range fanIn(sums) {
		res += sum
	}

	return res
}

// fanIn pattern
func fanIn(sums []<-chan int) <-chan int {
	res := make(chan int, len(sums))

	go func() {
		defer close(res)

		for i := 0; i < len(sums); i++ {
			select {
			case res <- <-sums[i]:
			default:
				log.Println("error: cannot receive from sum channel and send to res channel")
			}
		}
	}()

	return res
}

// generator pattern
func calculate(arr []int, wg *sync.WaitGroup) <-chan int {
	res := make(chan int, 1)

	go func() {
		sum := sumOfArray(arr)

		defer wg.Done()

		select {
		case res <- sum:
		default:
			log.Printf("error: cannot send sum to res channel")
		}
	}()

	return res
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

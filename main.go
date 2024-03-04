package main

func concurrentSum(arr []int) int {
	return 0
}

func naiveCuncurrentSum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {

}

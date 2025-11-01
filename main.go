package main

import "fmt"

func numWriter(out chan<- int, array [10]int) {
	for _, value := range array {
		out <- value
	}
	close(out)
}

func numProcessor(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * value
	}
	close(out)
}
func main() {

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	firstCh := make(chan int)
	secondCh := make(chan int)

	go numWriter(firstCh, array)

	go numProcessor(firstCh, secondCh)

	for value := range secondCh {
		fmt.Println(value)

	}
}

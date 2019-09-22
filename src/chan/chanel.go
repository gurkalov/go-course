package main

import "fmt"

func worker(ch chan int, len int) {
	for i := 0; i < len; i++ {
		fmt.Println("Read: ", <-ch)
	}
}

func main() {
	len := 3
	ch := make(chan int)

	go worker(ch, len)
	for i := 0; i < len; i++ {
		fmt.Println("Write: ", i)
		ch <- i
	}
}

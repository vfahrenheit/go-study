package main

import "fmt"

func main() {
	c := make(chan int,)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	c <- 3
	c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
}
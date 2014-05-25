package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	go func() {
		c1 <- "value 2"
		for i := 0; i < 100; i++ {
			fmt.Println("put value in chan", i)
		}
		fmt.Println("Done")
		// c1 <- "chan value"
	}()
	<-c1
	// <-c1
	// <-c1
	// a := <-c1

	// fmt.Println(a)
}

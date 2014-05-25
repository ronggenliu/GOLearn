package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	select {
	case msg := <-c1:
		fmt.Println("received ", msg)
	default:
		fmt.Println("not received")
	}
	<-c1
	select {
	case c1 <- "hello":
		fmt.Println("send msg ", "hello")
	default:
		fmt.Println("not send")
	}

	select {
	case msg1 := <-c1:
		fmt.Println("c1  receive msg ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 receive msg ", msg2)
	default:
		fmt.Println("both not receive")
	}
}

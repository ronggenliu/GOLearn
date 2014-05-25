package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "test1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "test2"
	}()

	//method 1:
	//-------------------------
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("defaut")
		}
	}

	// method 2:
	//--------------------------
	// msg1 := <-c1
	// fmt.Println("received", msg1)
	// msg2 := <-c2
	// fmt.Println("received", msg2)
}

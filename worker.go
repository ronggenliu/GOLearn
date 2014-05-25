package main

import (
	"fmt"
	"time"
)

var c chan bool

func worker() {
	fmt.Println("start working")
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
	c <- true
}

func ping(c1 chan<- string, s string) {
	c1 <- s
}

func pong(c1 <-chan string) {
	a := <-c1
	fmt.Println("get from chan: ", a)
}

func main() {
	c = make(chan bool)
	c1 := make(chan string, 1)
	ping(c1, "hello")
	pong(c1)
	go worker()
	<-c
}

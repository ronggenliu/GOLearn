package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	Try(func() {
		panic("foo")
	}, func(err interface{}) {
		fmt.Println(err)
	})
	var bytes []byte = make([]byte, 1024)
	f, err := os.Open("fmt.go")
	defer func() {
		fmt.Println("recover")
		if e := recover(); e != nil {
			fmt.Println("handle error: ", e)
		}
	}()
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := f.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()
	var er error
	_, er = f.Read(bytes)
	if er != nil {
		panic(er)
	}
BEGINFOR:
	for er == nil {
		fmt.Printf("%s", bytes)
		_, er = f.Read(bytes)
		if er != nil {
			log.Fatal("error: ", er)
			break BEGINFOR
		}
	}
	fmt.Printf("%s", bytes)
}

// try...catch...model
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

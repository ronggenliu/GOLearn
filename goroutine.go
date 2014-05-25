package main

import (
	"fmt"
)

func f1(arg string) {
	fmt.Printf("you input %s\n", arg)
}

func main() {
	go f1("lrg")

	go func() {
		fmt.Println("goroutine for anonymous method.")
	}()
	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}

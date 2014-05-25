package main

import (
	"fmt"
)

func main() {
	var j int = 10

	a := func() func() {
		var i int = 5
		return func() {
			i *= 2
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}()

	b := a()

	fmt.Printf("b=%v", b)

	j *= 2

	a()
}

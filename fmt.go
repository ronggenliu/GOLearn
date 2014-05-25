package main

import "fmt"

var (
	i1 = 1
	i2 = 2
)

type Person struct {
	name string
	age  int
}

func main() {
	me := Person{"ronggenliu", 32}

	fmt.Printf("%s %s %.5d %.*f \n", "s1", "s2", 100, 3, 9999.6666666)

	// %v 值输出。
	fmt.Printf("print me:%v\n", me)
	// %+V 值输出，带成员名
	fmt.Printf("print me:%+v\n", me)
	// %#v 值的GO表示
	fmt.Printf("print me:%#v", me)

	fmt.Printf("bool:%T\n", false)
	fmt.Printf("bool:%t\n", true)
	fmt.Printf("quote:%q\n", "ronggenliu")
	fmt.Printf("% :%%")
}

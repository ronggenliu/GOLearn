package main

import (
	"fmt"
	"github.com/astaxie/beego/utils"
	"path/filepath"
)

func main() {
	testfile := filepath.Join(".", "fmt.go")
	abolutepath, er := filepath.Abs(testfile)
	if er == nil {
		fmt.Println(abolutepath)
	}
	lines, err := utils.GrepFile("fmt", testfile)
	if err == nil {
		for _, line := range lines {
			fmt.Println(line)
		}
	} else {
		fmt.Println(err)
	}
}

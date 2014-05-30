package main

import (
	"fmt"
	"os"
	"time"
)

// 2014-05-30
func main() {
	// get host name.
	hostName, err := os.Hostname()
	if err == nil {
		fmt.Printf("Host Name: %q\n", hostName)
	}

	// get current absolute dir.
	curDir, er := os.Getwd()
	if er == nil {
		fmt.Printf("current dir: %q\n", curDir)
	}

	// get user id.
	userId := os.Geteuid()
	fmt.Printf("user id: %d\n", userId)

	f, e2 := os.Create("hello.txt")
	if e2 == nil {
		f.WriteString("hello-go, go, go...")
	}
	defer func() {
		f.Close()
		fmt.Println("close file hello.txt")
	}()
	b := make([]byte, 100)
	// from start.
	f.Seek(-3, 2)
	// from end.
	var n int
	n, _ = f.Read(b)
	fmt.Printf("read beyond 3 bytes from end offset: %s\n", b[0:n])
	f.Seek(3, 0)
	n, _ = f.Read(b)
	fmt.Printf("read beyond 3 from begin offset: %s\n", b[0:n])
	f.Seek(6, 0)
	f.Seek(6, 1)
	n, _ = f.Read(b)
	fmt.Printf("read beyond 6 from current offset: %s\n", b[:n])
	fstat, _ := f.Stat()
	fmt.Printf("file state: \n%+v\n", fstat)
	fmt.Printf("file state: \n%#v\n", fstat)

	fmt.Printf("%s is dir? %t, last modified time: %s\n", fstat.Name(), fstat.Mode().IsDir(), fstat.ModTime().Format(time.ANSIC))
}

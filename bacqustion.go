package main

import (
	"fmt"
	"time"
)

func capactiy(initBacNum int, tmins int) int {
	var capa int
	for i := 0; i < initBacNum; i++ {
		capa *= 2
	}
	return capa
}

func minsLapse(initBacNum int, capa int) int {
	var tmins int
	for i := 0; i < initBacNum; i++ {

	}
}

package main

import (
	"fmt"
	"time"
)

func POVERKA() {
	var TP1 string
	var TP2 string
	TP1 = "123456"
	TP2 = "654321"
	fmt.Printf(TP1)
	time.Sleep(5 * time.Second)
	fmt.Printf(TP2)
}

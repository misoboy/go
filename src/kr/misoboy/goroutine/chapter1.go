package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}

func main() {

	for i := 0; i < 100; i++ {
		go Hello(i)
	}

	fmt.Scanln()
	fmt.Println("Enter Exit")
}

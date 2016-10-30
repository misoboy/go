package main

import "fmt"
import "./goroutine"

func main(){

	for i := 0; i < 100 ; i++ {
		go goroutine.Hello(i)
	}

	fmt.Scanln()
	fmt.Println("Enter Exit")
}
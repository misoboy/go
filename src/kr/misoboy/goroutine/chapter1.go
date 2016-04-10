package goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello(n int){
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}
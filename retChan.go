package main

import (
	"fmt"
	"time"
)

func reverse(arr []int) <-chan int {
	r := make(chan int)

	go func() {
		defer close(r)
		// func() core (meaning, the operation to be completed)
		for i := len(arr) - 1; i >= 0; i-- {
			r <- arr[i]
			time.Sleep(time.Second * 10)
		}
	}()
	return r
}

func main() {
	potoc := reverse([]int{3, 4, 15})
	fmt.Println(<-potoc)
	fmt.Println(<-potoc)
	fmt.Println(<-potoc)
}

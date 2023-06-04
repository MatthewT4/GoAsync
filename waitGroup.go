package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func printLine(wg *sync.WaitGroup, num int) {
	iter := 1
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

PRINTER:
	for {
		select {
		case <-ctx.Done():
			break PRINTER
		default:
			fmt.Println(num, iter)
			iter += 1
		}
	}
	fmt.Printf("%v finish\n", num)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printLine(wg, i)
	}

	wg.Wait()
}

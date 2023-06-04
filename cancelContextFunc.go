package main

import (
	"context"
	"fmt"
	"time"
)

func Generator(ctx context.Context) {
	iter := 1
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println(iter)
			iter += 2
		}
	}
	fmt.Println("finish")
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	go Generator(ctx)
	fmt.Scanln()
}

package main

import "fmt"

func printAsync(cancelSignal <-chan interface{}, dataCh chan<- int) {
	val := 0
	for {
		select {
		case <-cancelSignal:
			return
		case dataCh <- val:
			val++
		}
	}
}

func main() {
	cancelCh := make(chan interface{}) // канал обмена
	dataCh := make(chan int, 3)
	go printAsync(cancelCh, dataCh)
	for i := range dataCh {
		fmt.Println(i)
		if i > 100 {
			cancelCh <- 1
			break
		}
	}
	close(dataCh)
}

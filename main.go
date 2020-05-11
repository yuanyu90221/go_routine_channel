package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)

	wg.Add(2)
	go func(ch <-chan int) {
		// i := <-ch
		// fmt.Println(i)
		// i = <-ch
		// fmt.Println(i)
		// for i := range ch {
		// 	fmt.Println(i)
		// }
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 100
		ch <- 101
		ch <- 102
		ch <- 103
		close(ch)
		// ch <- 105
		wg.Done()
	}(ch)

	wg.Wait()
}

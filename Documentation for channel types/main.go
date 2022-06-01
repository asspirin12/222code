package main

import (
	"fmt"
	"sync"
)

type MyFancyAdder struct {
	addBy int
}

func (a *MyFancyAdder) add(base int) int {
	return a.addBy + base
}

func dumbFibonacci(of int) int {
	if of < 2 {
		return of
	}
	return dumbFibonacci(of-2) + dumbFibonacci(of-1)
}

func dumbFibonacciInvoker(retchan chan<- string) {
	wg := sync.WaitGroup{}
	for i := 1; i < 40; i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			retchan <- fmt.Sprintf("fib(%v): %v\n", i, dumbFibonacci(i))
		}(i)
	}
	wg.Wait()
	close(retchan)
}

func main() {
	cf := make(chan string)
	dumbFibonacciInvoker(cf)
}

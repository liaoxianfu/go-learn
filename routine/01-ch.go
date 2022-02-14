package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printHello(i int) {
	defer wg.Done()
	fmt.Printf("hello %d\n", i)
}


func main(){
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printHello(i)
	}
	wg.Wait()
}


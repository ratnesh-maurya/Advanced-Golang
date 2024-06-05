package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	datachannel := make(chan int)

	go func() {
		datachannel <- 42
	}()
	data := <-datachannel
	fmt.Println(data)

	// main2()
	main3()

}

func main2() {
	datachannel := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			datachannel <- i
		}
		close(datachannel)
	}()

	for i := range datachannel {
		fmt.Println(i)
	}

}

func dowork() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(100)
}

func main3() {
	datachannel := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				datachannel <- dowork()
			}()
		}
		wg.Wait()
		close(datachannel)
	}()

	for n := range datachannel {
		fmt.Println(n)
	}

}

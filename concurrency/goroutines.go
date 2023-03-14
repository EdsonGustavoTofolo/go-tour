package concurrency

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func RunGoroutines() {
	go say("world")
	say("hello")
}

func RunChannels() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)

	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	fibonacciChan := make(chan int, 10)

	go fibonacci(cap(fibonacciChan), fibonacciChan)

	for i := range fibonacciChan {
		fmt.Println(i)
	}

	selectChannel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-selectChannel)
		}
		quit <- 0
	}()
	fibonacciSelect(selectChannel, quit)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

package multithreading

import (
	"fmt"
	"sync"
	"time"
)

func RunChannels() {
	channel := make(chan string)

	go func() {
		channel <- "Hello World"
		channel <- "Hello World 2"
	}()

	msg := <-channel

	fmt.Println(msg, <-channel)
}

func RunChannelsRange() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func reader(ch chan int) {
	for value := range ch {
		fmt.Printf("Received %d\n", value)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // avoid deadlock
}

func RunChannelsRangeWithWaitGroups() {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)

	go func() {
		for value := range ch {
			fmt.Printf("Received %d\n", value)
			wg.Done()
		}
	}()

	wg.Wait()
}

func RunChannelsDirections() {
	hello := make(chan string)

	go receive("Hello", hello)

	fmt.Println(read(hello))
}

func receive(name string, hello chan<- string) {
	hello <- name
}

func read(data <-chan string) string {
	return <-data
}

func RunChannelsWithSelect() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		println(msg1)
	case msg2 := <-c2:
		println(msg2)
	case <-time.After(time.Second * 3):
		println("timeout")
	}
}

func RunChannelsWithBuffer() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)

}

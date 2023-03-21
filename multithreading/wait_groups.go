package multithreading

import (
	"fmt"
	"sync"
	"time"
)

func task2(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func RunGoRoutinesWithWaitGroups() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)

	go task2("A", &waitGroup)
	go task2("B", &waitGroup)
	go func(name string) {
		task2(name, &waitGroup)
	}("C")

	waitGroup.Wait()
}

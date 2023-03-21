package multithreading

import (
	"fmt"
	"time"
)

func task1(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func RunGoRoutines() {
	go task1("A")
	go task1("B")
	go func(name string) {
		task1(name)
	}("C")

	time.Sleep(10 * time.Second)
}

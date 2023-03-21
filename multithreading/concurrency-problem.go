package multithreading

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var number uint64 = 0

/*
RunConcurrencyRealProblem

How to check the problem:
In terminal enter:
> ab -n 10000 -c 100 http://localhost:3000/
*/
func RunConcurrencyRealProblem() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		number++
		time.Sleep(300 * time.Millisecond)
		writer.Write([]byte(fmt.Sprintf("You are the visitant number %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

func RunConcurrencyRealProblemMutexSolver() {
	m := sync.Mutex{}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		m.Lock()
		number++
		m.Unlock()

		time.Sleep(300 * time.Millisecond)

		writer.Write([]byte(fmt.Sprintf("You are the visitant number %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

func RunConcurrencyRealProblemAtomicSolver() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		atomic.AddUint64(&number, 1)

		time.Sleep(300 * time.Millisecond)

		writer.Write([]byte(fmt.Sprintf("You are the visitant number %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

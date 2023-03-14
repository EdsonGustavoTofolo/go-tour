package moretypes

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func Matrix() {
	initializingMatrix()

	var a [2]string
	a[0] = "Edson"
	a[1] = "Tofolo"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11}

	fmt.Println(primes)
}

func initializingMatrix() {
	primes := [6]int{2, 3, 5, 7, 11}
	var a [2]string

	fmt.Println(primes, a)
}

func Slices() {
	initializingSlices()

	addingItemOnSlice()

	primes := [6]int{2, 3, 5, 7, 11}

	var s []int = primes[1:4]

	fmt.Println(s)

	q := []int{2, 4, 6, 8, 10}
	fmt.Println(q)

	s = []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
}

func addingItemOnSlice() {
	a := make([]int, 3)
	a = append(a, 1)
	a = append(a, 2)

	printSlice(a)
}

func initializingSlices() {
	primes := [6]int{2, 3, 5, 7, 11}
	var a []int = primes[1:4]

	s := []int{2, 3, 5, 7, 11, 13}

	// init slice with 5 items with value 0 -> length 5 capacity 5, vide len(), cap()
	b := make([]int, 5)

	// init empty slice with capacity 5 -> length 0 capacity 5. vide len(), cap()
	c := make([]int, 0, 5)

	fmt.Println(a, s, b, c)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Maps() {
	m := initializingMap()

	m["a"] = 1

	fmt.Println(m)
	fmt.Println(m["a"])

	mutationMap()
}

func mutationMap() {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println(m)

	delete(m, "a")

	fmt.Println(m)

	v, ok := m["a"]
	fmt.Println("THe value:", v, "Present?", ok)
}

func initializingMap() map[string]int {
	var m map[string]int

	m = make(map[string]int)

	x := map[string]Vertex{
		"A": {Lat: 0, Long: 1},
		"B": {Lat: 0, Long: 1},
	}

	fmt.Println(x)

	return m
}

package methods

import (
	"fmt"
	"math"
)

type Number interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v *Vertex) Abs() float64 {
	v.X = 10
	v.Y = 5
	return math.Sqrt(v.X * v.Y)
}

func Abs(v *Vertex) float64 {
	return v.Abs()
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Interfaces() {
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}

func Do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

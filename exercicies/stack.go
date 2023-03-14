package exercicies

import (
	"errors"
	"fmt"
	"github.com/gammazero/deque"
)

type Numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Stack[T Numbers] struct {
	values    deque.Deque[T]
	minValues deque.Deque[T]
}

func (f *Stack[T]) Push(value T) T {
	if f.minValues.Len() == 0 {
		f.minValues.PushBack(value)
	} else {
		lastValue := f.minValues.At(f.values.Len() - 1)
		if value < lastValue {
			f.minValues.PushBack(value)
		} else {
			f.minValues.PushBack(lastValue)
		}
	}

	f.values.PushBack(value)

	return value
}

func (f *Stack[T]) Pop() (T, error) {
	if f.values.Len() == 0 {
		return -1, errors.New("stack is empty")
	}

	f.minValues.PopBack()

	return f.values.PopBack(), nil
}

func (f *Stack[T]) GetMin() (T, error) {
	if f.values.Len() == 0 {
		return -1, errors.New("stack is empty")
	}
	return f.minValues.At(f.minValues.Len() - 1), nil
}

func RunStack() {
	stack := Stack[int]{}

	stack.Push(6)
	stack.Push(5)
	stack.Push(10)
	stack.Push(8)
	stack.Push(5)
	stack.Push(2)

	min, _ := stack.GetMin()

	fmt.Println("Min:", min)

	popped, _ := stack.Pop()

	fmt.Println("Popped:", popped)

	min, _ = stack.GetMin()

	fmt.Println("Min:", min)

	popped, _ = stack.Pop()

	fmt.Println("Popped:", popped)

	min, _ = stack.GetMin()

	fmt.Println("Min:", min)

	popped, _ = stack.Pop()

	fmt.Println("Popped:", popped)

	min, _ = stack.GetMin()

	fmt.Println("Min:", min)

	popped, _ = stack.Pop()

	fmt.Println("Popped:", popped)

	min, _ = stack.GetMin()

	fmt.Println("Min:", min)

	popped, _ = stack.Pop()

	fmt.Println("Popped:", popped)

	min, _ = stack.GetMin()

	fmt.Println("Min:", min)

	popped, _ = stack.Pop()

	fmt.Println("Popped:", popped)

	min, err := stack.GetMin()

	if err != nil {
		fmt.Println("-->", err.Error())
	}
}

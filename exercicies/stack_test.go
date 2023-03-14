package exercicies

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(6)

	min, err := stack.GetMin()

	assert.NoError(t, err)
	assert.Equal(t, 6, min)
}

func TestTwo(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(6)
	stack.Push(5)
	stack.Push(10)

	min, err := stack.GetMin()

	assert.NoError(t, err)
	assert.Equal(t, 5, min)
}

func TestThree(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(6)
	stack.Push(5)
	stack.Push(5)

	min, err := stack.GetMin()

	assert.NoError(t, err)
	assert.Equal(t, 5, min)
}

func TestFour(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(6)
	stack.Push(5)
	stack.Push(5)
	stack.Push(10)

	popped, err := stack.Pop()

	assert.NoError(t, err)
	assert.Equal(t, 10, popped)

	min, err := stack.GetMin()

	assert.NoError(t, err)
	assert.Equal(t, 5, min)
}

func TestFive(t *testing.T) {
	stack := Stack[int]{}

	popped, err := stack.Pop()

	assert.Error(t, err, "stack is empty")
	assert.Equal(t, -1, popped)
}

func TestSix(t *testing.T) {
	stack := Stack[int]{}

	min, err := stack.GetMin()

	assert.Error(t, err, "stack is empty")
	assert.Equal(t, -1, min)
}

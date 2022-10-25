package stack

import (
	"errors"
	"sync"
)

var EmptyStack = errors.New("empty Stack")

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return make([]T, 0)
}

func (st *Stack[T]) zeroT() T {
	var zero T
	return zero
}

func (st *Stack[T]) Push(value T) {
	*st = append(*st, value)
}

func (st *Stack[T]) Empty() bool {
	return len(*st) == 0
}

func (st *Stack[T]) Pop() (T, error) {
	if st.Empty() {
		return st.zeroT(), EmptyStack
	}
	size := len(*st)
	top := (*st)[size-1]
	*st = (*st)[:size-1]
	return top, nil
}

func (st *Stack[T]) Top() (T, error) {
	if st.Empty() {
		return st.zeroT(), EmptyStack
	}
	return (*st)[0], nil
}

type AsyncStack[T any] struct {
	stack Stack[T]
	lock  sync.Mutex
}

func NewAsyncStack[T any]() *AsyncStack[T] {
	return &AsyncStack[T]{
		lock:  sync.Mutex{},
		stack: NewStack[T]()}
}

func (st *AsyncStack[T]) Push(value T) {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.stack.Push(value)
}

func (st *AsyncStack[T]) Empty() bool {
	return st.stack.Empty()
}

func (st *AsyncStack[T]) Pop() (T, error) {
	st.lock.Lock()
	defer st.lock.Unlock()

	return st.stack.Pop()
}

func (st *AsyncStack[T]) Top() (T, error) {
	return st.stack.Top()
}

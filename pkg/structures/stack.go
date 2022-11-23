package structures

import (
	"errors"
	"sync"
)

var ErrorEmptyStack = errors.New("empty Stack")

type Stack[T any] struct {
	deq Dequeue[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		deq: NewDequeue[T](),
	}
}

func (st *Stack[T]) Push(value T) {
	st.deq.PushBack(value)
}

func (st *Stack[T]) Empty() bool {
	return st.deq.Empty()
}

func (st *Stack[T]) Pop() (T, error) {
	res, err := st.deq.PopBack()
	if err != nil {
		return res, ErrorEmptyStack
	}

	return res, nil
}

func (st *Stack[T]) MustPop() T {
	res, err := st.Pop()
	if err != nil {
		panic(err)
	}

	return res
}

func (st *Stack[T]) Top() (T, error) {
	return st.deq.Back()
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

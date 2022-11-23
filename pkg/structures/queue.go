package structures

import (
	"errors"
	"sync"
)

var ErrorEmptyQueue = errors.New("empty queue")

type Queue[T any] struct {
	deq Dequeue[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		deq: NewDequeue[T](),
	}
}

func (st *Queue[T]) Push(value T) {
	st.deq.PushBack(value)
}

func (st *Queue[T]) Empty() bool {
	return st.deq.Empty()
}

func (st *Queue[T]) Pop() (T, error) {
	res, err := st.deq.PopFront()
	if err != nil {
		return res, ErrorEmptyQueue
	}

	return res, nil
}

func (st *Queue[T]) MustPop() T {
	res, err := st.Pop()
	if err != nil {
		panic(err)
	}

	return res
}

func (st *Queue[T]) Front() (T, error) {
	return st.deq.Front()
}

type AsyncQueue[T any] struct {
	queue Queue[T]
	lock  sync.Mutex
}

func NewAsyncQueue[T any]() *AsyncQueue[T] {
	return &AsyncQueue[T]{
		lock:  sync.Mutex{},
		queue: NewQueue[T]()}
}

func (st *AsyncQueue[T]) Push(value T) {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.queue.Push(value)
}

func (st *AsyncQueue[T]) Empty() bool {
	return st.queue.Empty()
}

func (st *AsyncQueue[T]) Pop() (T, error) {
	st.lock.Lock()
	defer st.lock.Unlock()

	return st.queue.Pop()
}

func (st *AsyncQueue[T]) Front() (T, error) {
	return st.queue.Front()
}

package structures

import (
	"errors"
	"sync"

	"golang.org/x/exp/slices"
)

var ErrorEmptyDequeue = errors.New("empty dequeue")

type Dequeue[T any] []T

func NewDequeue[T any]() Dequeue[T] {
	return make([]T, 0)
}

func (st *Dequeue[T]) zeroT() T {
	var zero T

	return zero
}

func (st *Dequeue[T]) PushBack(value T) {
	*st = append(*st, value)
}

func (st *Dequeue[T]) PushFront(value T) {
	*st = slices.Insert(*st, 0, value)
}

func (st *Dequeue[T]) Empty() bool {
	return len(*st) == 0
}

func (st *Dequeue[T]) PopBack() (T, error) {
	if st.Empty() {
		return st.zeroT(), ErrorEmptyDequeue
	}

	size := len(*st)
	back := (*st)[size-1]

	*st = (*st)[:size-1]

	return back, nil
}

func (st *Dequeue[T]) MustPopBack() T {
	res, err := st.PopBack()
	if err != nil {
		panic(err)
	}

	return res
}

func (st *Dequeue[T]) Back() (T, error) {
	if st.Empty() {
		return st.zeroT(), ErrorEmptyDequeue
	}

	return (*st)[len(*st)-1], nil
}

func (st *Dequeue[T]) Front() (T, error) {
	if st.Empty() {
		return st.zeroT(), ErrorEmptyDequeue
	}

	return (*st)[0], nil
}

func (st *Dequeue[T]) PopFront() (T, error) {
	if st.Empty() {
		return st.zeroT(), ErrorEmptyDequeue
	}

	front := (*st)[0]
	*st = (*st)[1:]

	return front, nil
}

func (st *Dequeue[T]) MustPopFront() T {
	res, err := st.PopFront()
	if err != nil {
		panic(err)
	}

	return res
}

type AsyncDequeue[T any] struct {
	dequeue Dequeue[T]
	lock    sync.Mutex
}

func NewAsyncDequeue[T any]() *AsyncDequeue[T] {
	return &AsyncDequeue[T]{
		lock:    sync.Mutex{},
		dequeue: NewDequeue[T](),
	}
}

func (st *AsyncDequeue[T]) PushBack(value T) {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.dequeue.PushBack(value)
}

func (st *AsyncDequeue[T]) PushFront(value T) {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.dequeue.PushFront(value)
}

func (st *AsyncDequeue[T]) Empty() bool {
	return st.dequeue.Empty()
}

func (st *AsyncDequeue[T]) PopBack() (T, error) {
	st.lock.Lock()
	defer st.lock.Unlock()

	return st.dequeue.PopBack()
}

func (st *AsyncDequeue[T]) MustPopBack() T {
	st.lock.Lock()
	defer st.lock.Unlock()

	return st.dequeue.MustPopBack()
}

func (st *AsyncDequeue[T]) Back() (T, error) {
	return st.dequeue.Back()
}

func (st *AsyncDequeue[T]) Front() (T, error) {
	return st.dequeue.Front()
}

func (st *AsyncDequeue[T]) PopFront() (T, error) {
	st.lock.Lock()
	defer st.lock.Unlock()

	return st.dequeue.PopFront()
}

func (st *AsyncDequeue[T]) MustPopFront() T {
	st.lock.Lock()
	defer st.lock.Unlock()

	return st.dequeue.MustPopFront()
}

package graph

type ValueNode[T any] struct {
    Value T
}

type Node[T any, Label comparable] struct {
    ValueNode[T]
    next map[Label]*Node[T, Label]
}

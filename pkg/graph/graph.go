package graph

import "github.com/ThCompiler/go_game_constractor/pkg/structures"

type Graph[T any, Label comparable] struct {
    graph map[Label]*Node[T, Label]
}

type VertexInfo[T any, Label comparable] struct {
    Label Label
    Value T
}

type EdgeInfo[Label comparable] struct {
    VertexFrom Label
    VertexTo   Label
}

type Visitor[T any] func(value *ValueNode[T]) (stop bool)

func NewGraph[T any, Label comparable](vertexes []VertexInfo[T, Label]) *Graph[T, Label] {
    graph := &Graph[T, Label]{}
    graph.SetVertexes(vertexes)
    return graph
}

func (graph *Graph[T, Label]) SetVertexes(vertexes []VertexInfo[T, Label]) {
    graph.graph = make(map[Label]*Node[T, Label])
    for _, vertex := range vertexes {
        graph.AddVertex(vertex.Label, vertex.Value)
    }
}

func (graph *Graph[T, Label]) AddVertex(label Label, Value T) {
    graph.graph[label] = &Node[T, Label]{
        ValueNode: ValueNode[T]{Value},
        next:      make(map[Label]*Node[T, Label]),
    }
}

func (graph *Graph[T, Label]) DeleteVertex(label Label) {
    delete(graph.graph, label)
}

func (graph *Graph[T, Label]) AddOrientedEdges(edges ...EdgeInfo[Label]) (err error) {
    for _, edge := range edges {
        err = graph.AddOrientedEdge(edge.VertexFrom, edge.VertexTo)
        if err != nil {
            break
        }
    }

    return err
}

func (graph *Graph[T, Label]) AddUndirectedEdges(edges ...EdgeInfo[Label]) (err error) {
    for _, edge := range edges {
        err = graph.AddUndirectedEdge(edge.VertexFrom, edge.VertexTo)
        if err != nil {
            break
        }
    }

    return err
}

func (graph *Graph[T, Label]) AddOrientedEdge(vertexFrom Label, vertexTo Label) error {
    if _, is := graph.graph[vertexFrom]; !is {
        return ErrorNotFoundVertex
    }

    if _, is := graph.graph[vertexTo]; !is {
        return ErrorNotFoundVertex
    }

    graph.graph[vertexFrom].next[vertexTo] = graph.graph[vertexTo]
    return nil
}

func (graph *Graph[T, Label]) AddUndirectedEdge(vertexFirst Label, vertexSecond Label) error {
    if err := graph.AddOrientedEdge(vertexFirst, vertexSecond); err != nil {
        return err
    }

    return graph.AddOrientedEdge(vertexSecond, vertexFirst)
}

func (graph *Graph[T, Label]) DeleteOrientedEdge(vertexFrom Label, vertexTo Label) error {
    if _, is := graph.graph[vertexFrom]; !is {
        return ErrorNotFoundVertex
    }

    if _, is := graph.graph[vertexTo]; !is {
        return ErrorNotFoundVertex
    }

    delete(graph.graph[vertexFrom].next, vertexTo)
    return nil
}

func (graph *Graph[T, Label]) DeleteUndirectedEdge(vertexFirst Label, vertexSecond Label) error {
    if err := graph.DeleteOrientedEdge(vertexFirst, vertexSecond); err != nil {
        return err
    }

    return graph.DeleteOrientedEdge(vertexSecond, vertexFirst)
}

func (graph *Graph[T, Label]) DFS(startVertex Label, visitor Visitor[T]) {
    dfsStack := structures.NewStack[*Node[T, Label]]()
    visited := graph.createVisited()

    dfsStack.Push(graph.graph[startVertex])

    for !dfsStack.Empty() {
        currentVertex, _ := dfsStack.Pop()

        if visitor(&currentVertex.ValueNode) {
            break
        }

        for label, vertex := range currentVertex.next {
            if !visited[label] {
                visited[label] = true
                dfsStack.Push(vertex)
            }
        }
    }
}

func (graph *Graph[T, Label]) BFS(startVertex Label, visitor Visitor[T]) {
    bfsQueue := structures.NewQueue[*Node[T, Label]]()
    visited := graph.createVisited()

    bfsQueue.Push(graph.graph[startVertex])

    for !bfsQueue.Empty() {
        currentVertex, _ := bfsQueue.Pop()

        if visitor(&currentVertex.ValueNode) {
            break
        }

        for label, vertex := range currentVertex.next {
            if !visited[label] {
                visited[label] = true
                bfsQueue.Push(vertex)
            }
        }
    }
}

func (graph *Graph[T, Label]) createVisited() map[Label]bool {
    visited := make(map[Label]bool)
    for vertex := range graph.graph {
        visited[vertex] = false
    }
    return visited
}

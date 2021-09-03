package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

func (v Vertex) Contains(n int) bool {
	for _, v := range v.Adjacent {
		if n == v.Key {
			return true
		}
	}

	return false
}

func (g *Graph) AddVertex(key int) error {
	if g.Contains(key) {
		return errors.New("key already exists in the graph")
	}

	g.Vertices = append(g.Vertices, &Vertex{Key: key})

	return nil
}

func (g *Graph) AddEdge(from, to int) error {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex == nil || toVertex == nil {
		return fmt.Errorf("invalid edge (%v -> %v)", from, to)
	}

	if fromVertex.Contains(to) {
		return fmt.Errorf("edge already exists (%v -> %v)", from, to)
	}

	fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)

	return nil
}

func (g Graph) Contains(n int) bool {
	for _, v := range g.Vertices {
		if n == v.Key {
			return true
		}
	}

	return false
}

func (g Graph) GetVertex(key int) *Vertex {
	for i, v := range g.Vertices {
		if v.Key == key {
			return g.Vertices[i]
		}
	}

	return nil
}

func (g Graph) PrintData() {
	for _, v := range g.Vertices {
		fmt.Println(v.Key)

		for _, v := range v.Adjacent {
			fmt.Println(v.Key)
		}
	}
}

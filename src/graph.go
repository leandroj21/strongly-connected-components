package src

import (
	"fmt"
	"time"
)

var (
	t     int
	stack []int
	arbol []int
)

type intTuple [2]int

type Node struct {
	Label     int
	Visited   int
	Previous  int
	Neighbors []int
}

type Graph struct {
	Nodes []*Node
}

func (g *Graph) CreateGraph(lsls []intTuple, rev bool, start time.Time) {

}

func (g *Graph) Display() {
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		near := g.Nodes[i]
		if near == nil {
			continue
		}
		s += fmt.Sprintf(" %8d --> ", near.Label)
		for _, j := range near.Neighbors {
			s += fmt.Sprintf(" %8d| ", g.Nodes[j].Label)
		}
		fmt.Println(s)
		s = ""
	}
}

func (g *Graph) Dfs(rev bool) {

}

// Operations

func DfsVisit(g *Graph, i int, u *Node, rev bool) {

}

func ListOrder(pg *Graph) {

}

func FindLeader(g *Graph) {

}

func ListTree(pg *Graph) {

}
